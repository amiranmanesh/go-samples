package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func SetCustomValidations() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("db_unique", dbUnique)
		if err != nil {
			logrus.Error(err)
		}

		err2 := v.RegisterValidation("db_exists", dbExists)
		if err2 != nil {
			logrus.Error(err)
		}

		err3 := v.RegisterValidation("phone", phoneValidaion)
		if err3 != nil {
			logrus.Error(err)
		}
	}

}
