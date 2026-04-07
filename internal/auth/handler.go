package auth

import (
	"encoding/json"
	"net/http"
	"todo/internal/apperr"
	"todo/internal/user"
	"todo/internal/utils"

	"github.com/go-playground/validator/v10"
)


var validate = validator.New();

func init() {
	validate.RegisterValidation("password", utils.PasswordValidator);
}

// - Throws BadJSONBodyError
// - Throws BadJSONBodyError
// - Throws UserFound
// - Throws any, Database error.
// - Throws any, The rand function can read.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var dto RegisterDto;
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		panic(apperr.BadJSONBodyError);
	}

	if err := validate.Struct(dto); err != nil {
		panic(apperr.BodyValidationError);
	}

	newUser := user.UserCreate(h.db, &user.User{
		Name: dto.Name,
		Age:  dto.Age,
		Email: dto.Email,
		Password: HashPassword(dto.Password),
	});

	utils.WriteResponse(w, http.StatusCreated, newUser);
}
