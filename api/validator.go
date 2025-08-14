package api

import (
	"github.com/devsirose/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var isValidCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
