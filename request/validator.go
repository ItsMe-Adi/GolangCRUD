package request

import (
	"playground/util"

	"github.com/go-playground/validator/v10"
)

var ValidRole validator.Func = func(fl validator.FieldLevel) bool {
	if role, ok := fl.Field().Interface().(string); ok {
		return util.IsValidRole(role)
	}
	return false
}
