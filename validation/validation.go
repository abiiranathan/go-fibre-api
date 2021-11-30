package validation

import (
	"github.com/abiiranathan/goclinic/errors"
	"github.com/go-playground/validator"
)

func ValidateModel(model interface{}) []*errors.ErrorResponse {
	var errorList []*errors.ErrorResponse
	validate := validator.New()
	err := validate.Struct(model)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorResp errors.ErrorResponse

			errorResp.FailedField = err.StructNamespace()
			errorResp.Tag = err.Tag()
			errorResp.Value = err.Param()

			errorList = append(errorList, &errorResp)
		}
	}

	return errorList
}
