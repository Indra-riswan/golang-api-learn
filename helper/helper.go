package helper

import "strings"

type Respon struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildRespons(status bool, message string, data interface{}) Respon {
	res := Respon{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorRespons(message string, errors string, data interface{}) Respon {
	errSplit := strings.Split(errors, "\n")

	res := Respon{
		Status:  false,
		Message: message,
		Errors:  errSplit,
		Data:    data,
	}
	return res
}
