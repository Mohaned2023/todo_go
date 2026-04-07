package auth

import (
	"encoding/json"
	"net/http"
	"todo/internal/apperr"
	"todo/internal/user"
	"todo/internal/utils"
	"todo/pkg/logger"

	"github.com/go-playground/validator/v10"
)


var validate = validator.New();

func init() {
	validate.RegisterValidation("password", utils.PasswordValidator);
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var dto RegisterDto;
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		panic(apperr.BadJSONBodyError);
	}

	if err := validate.Struct(dto); err != nil {
		panic(apperr.BodyValidationError);
	}

	newUser, err := user.UserCreate(h.db, &user.User{
		Name: dto.Name,
		Age:  dto.Age,
		Email: dto.Email,
		Password: dto.Password,
	});
	if err != nil {
		logger.Err(err.Error());
		return;
	}

	utils.WriteResponse(w, http.StatusCreated, newUser);
}
