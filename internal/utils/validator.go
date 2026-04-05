package utils

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String();
	var (
		upper,
		lower,
		number bool
	);
	for _, c := range password {
		switch {
			case unicode.IsUpper(c):
				upper = true
			case unicode.IsLower(c):
				lower = true
			case unicode.IsNumber(c):
				number = true
		}
	}
	return upper && lower && number;
}
