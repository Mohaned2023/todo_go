package auth

import (
	"encoding/json"
	"net/http"
	"todo/internal/apperr"
	"todo/internal/session"
	"todo/internal/user"
	"todo/internal/utils"

	"github.com/go-playground/validator/v10"
)


var validate = validator.New();

func init() {
	validate.RegisterValidation("password", utils.PasswordValidator);
}

// - Throws BadJSONBodyError
// - Throws BodyValidationError
// - Throws UserFound
// - Throws any, Database error.
// - Throws any, The rand function can read.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var dto RegisterDto;
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		panic(apperr.Exception{
			Type: apperr.BadJSONBodyError,
			More: nil,
		});
	}

	if err := validate.Struct(dto); err != nil {
		panic(apperr.Exception{
			Type: apperr.BodyValidationError,
			More: nil,
		});
	}

	newUser := user.UserCreate(h.db, &user.User{
		Name: dto.Name,
		Age:  dto.Age,
		Email: dto.Email,
		Password: HashPassword(dto.Password),
	});

	utils.WriteResponse(w, http.StatusCreated, newUser);
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	rSid, err := r.Cookie("sid")
	if err == nil {
		u, err := session.GetAndUpdateTTL(r.Context(), h.redisClient, rSid.Value)
		if err == nil {
			utils.WriteResponse(w, http.StatusOK, u)
			return
		}
	}

	var dto LoginDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		panic(apperr.Exception{
			Type: apperr.BadJSONBodyError,
			More: nil,
		})
	}

	if err := validate.Struct(dto); err != nil {
		panic(apperr.Exception{
			Type: apperr.BodyValidationError,
			More: nil,
		})
	}

	user := user.GetUser(r.Context(), h.db, dto.Email)
	if !ComparePassword(dto.Password, user.Password) {
		panic(apperr.Exception{
			Type: apperr.UnauthorizedUser,
			More: nil,
		})
	}

	sid := session.MakeSession(r.Context(), h.redisClient, user)
	utils.WriteCookie(w, "sid", sid)
	utils.WriteResponse(w, http.StatusOK, user)
}
