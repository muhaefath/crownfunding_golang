package helper

import "github.com/go-playground/validator/v10"

type Respone struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIRespone(message string, code int, status string, data interface{}) Respone {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	Jsonresponse := Respone{
		Meta: meta,
		Data: data,
	}

	return Jsonresponse
}

func FormatValidationError(err error) []string {

	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
