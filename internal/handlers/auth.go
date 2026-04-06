package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"todo/internal/services"
	"todo/internal/storage"
	"todo/internal/types"
	"todo/internal/utils"
	"todo/pkg/logger"
);

var validate = validator.New();

func init() {
	validate.RegisterValidation("password", utils.PasswordValidator);
}

func AuthRegister(w http.ResponseWriter, req *http.Request) {
	var dto types.RegisterDto;
	if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
		panic(types.BadJSONBodyError);
	}

	if err := validate.Struct(dto); err != nil {
		panic(types.BodyValidationError);
	}

	newUser, err := services.UserCreate(storage.DBConn, &types.User{
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
