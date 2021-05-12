package validation

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func EmailValidate(email string) bool {
	//regular expression for email
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

func GetValidationErrors(err error) []map[string]string {
	validatorErrors := make([]map[string]string, 0)

	switch err.(type) {
	case validator.ValidationErrors:
		{
			for _, fieldErr := range err.(validator.ValidationErrors) {
				item := map[string]string{}
				item[fieldErr.Field()] = fieldErr.Tag()
				validatorErrors = append(validatorErrors, item)
			}

		}
	case *json.UnmarshalTypeError:
		{
			item := map[string]string{}
			item["json_bind_error"] = err.Error()
			validatorErrors = append(validatorErrors, item)

		}
	default:
		{
			item := map[string]string{}
			item["unknown_error"] = err.Error()
			validatorErrors = append(validatorErrors, item)
		}
	}
	return validatorErrors
}

var phoneValidaion validator.Func = func(fl validator.FieldLevel) bool {

	value := fl.Field().String()
	return PhoneRegex(value)
}

func PhoneRegex(phone string) bool {
	matched, err := regexp.MatchString(`^09\d{9}$`, phone)
	if err != nil {
		return false
	}
	return matched
}
