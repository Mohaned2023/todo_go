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
)

type AppError struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (a *AppError) Map(e AppErrorTypes) {
	a.Type = fmt.Sprintf("TD%0*d", 4, e);
	switch(e) {
		case InteralServerError:
			a.Status = http.StatusInternalServerError
			a.Message = "Internal Server Error"
		case BadJSONBodyError:
			a.Status = http.StatusBadRequest
			a.Message = "Bad json body"
		case BodyValidationError:
			a.Status = http.StatusBadRequest
			a.Message = "Invalid body fields"
	}
}
