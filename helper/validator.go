package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TrasnsalateErrorMessage(err error) map[string]string {
	errorMap := make(map[string]string)

	if validationError, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationError {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorMap[field] = fmt.Sprintf("%s is requaired", field)
			case "email":
				errorMap[field] = "Invalid Email Format"
			case "unique":
				errorMap[field] = fmt.Sprintf("%s is already exsist", field)
			case "min":
				errorMap[field] = fmt.Sprintf("%s must be at lesat %s characters", field, fieldError.Param())
			case "max":
				errorMap[field] = fmt.Sprintf("%s must be at lesat %s characters", field, fieldError.Param())
			case "numeric":
				errorMap[field] = fmt.Sprintf("%s is already exsist", field)
			default:
				errorMap[field] = "Invalid Value"
			}
		}
	}

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username"){
				errorMap["Username"] = "Username Already Exsists"
			}
			if strings.Contains(err.Error(), "email") {
				errorMap["Email"] = "Email Already Exsists"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorMap["Error"] = "Record Not Found"
		}
	}
	return errorMap
}

func isDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplcate Entry")
}
