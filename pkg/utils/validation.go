package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func BindAndValidate(r *http.Request, req interface{}) error {
	body := json.NewDecoder(r.Body)
	//body.DisallowUnknownFields()

	if err := body.Decode(req); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return fmt.Errorf("invalid validation error: %v", err)
		}

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				var errMsg string
				switch e.Tag() {
				case "required":
					errMsg = fmt.Sprintf("field %s is required", e.Field())
				case "datetime":
					errMsg = fmt.Sprintf("field %s must be in the format yyyy-MM-dd HH:mm:ss", e.Field())
				case "min":
					errMsg = fmt.Sprintf("field %s must be at least %s", e.Field(), e.Param())
				case "max":
					errMsg = fmt.Sprintf("field %s must be at most %s", e.Field(), e.Param())
				default:
					errMsg = fmt.Sprintf("field %s failed validation for tag %s", e.Field(), e.Tag())
				}

				return fmt.Errorf(errMsg)
			}
		}

		return err
	}

	return nil
}
func datetimeValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02 15:04:05", fl.Field().String())
	return err == nil
}
