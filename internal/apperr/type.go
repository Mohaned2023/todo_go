package apperr

import (
	"fmt"
	"net/http"
)

type AppErrorTypes int
const (
	InteralServerError AppErrorTypes = iota
	BadJSONBodyError
	BodyValidationError
	UserFound
	UserNotFound
	UnauthorizedUser
)

type Exception struct {
	Type AppErrorTypes
	More *map[string] any
}

type AppError struct {
	Status  int              `json:"-"`
	Message string           `json:"message"`
	Type    string           `json:"type"`
	More    *map[string] any `json:"more"`
}

func (a *AppError) Map(e Exception) {
	a.Type = fmt.Sprintf("TD%0*d", 4, e.Type);
	a.More = e.More;
	switch(e.Type) {
		case InteralServerError:
			a.Status = http.StatusInternalServerError
			a.Message = "Internal Server Error"
		case BadJSONBodyError:
			a.Status = http.StatusBadRequest
			a.Message = "Bad json body"
		case BodyValidationError:
			a.Status = http.StatusBadRequest
			a.Message = "Invalid body fields"
		case UserFound:
			a.Status = http.StatusFound
			a.Message = "User already registerd"
		case UserNotFound:
			a.Status = http.StatusNotFound
			a.Message = "User not found"
		case UnauthorizedUser:
			a.Status = http.StatusUnauthorized
			a.Message = "You are unauthorized"
	}
}
