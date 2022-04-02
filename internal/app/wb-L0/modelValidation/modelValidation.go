package modelValidation

import (
	"github.com/go-playground/validator"
	"wb-L0/internal/app/wb-L0/logger"
	"wb-L0/internal/app/wb-L0/storage"
)

func Validate(model *storage.ModelJSON) bool {
	v := validator.New()

	err := v.Struct(model)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			logger.Log.Error(e)
		}
		return true
	}
	return false
}
