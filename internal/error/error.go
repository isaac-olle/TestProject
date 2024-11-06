package error

import (
	"strconv"
)

type HttpError struct {
	httpCode    int
	description string
}

func NewHttpError(httpCode int, description string) *HttpError {
	return &HttpError{httpCode, description}
}

func (this *HttpError) Error() string {
	return "Request Failed. Http Code Error:" + strconv.Itoa(this.httpCode) + ", Description:" + this.description
}

func (this *HttpError) HttpCode() int {
	return this.httpCode
}

func (this *HttpError) Description() string {
	return this.description
}

type BadRequestHttpError struct {
	*HttpError
}

func NewBadRequestHttpError(description string) *BadRequestHttpError {
	return &BadRequestHttpError{NewHttpError(400, description)}
}

type NotFoundHttpError struct {
	*HttpError
}

func NewNotFoundHttpError(description string) *NotFoundHttpError {
	return &NotFoundHttpError{NewHttpError(404, description)}
}

type InternalServerError struct {
	*HttpError
}

func NewInternalServerError(description string) *InternalServerError {
	return &InternalServerError{NewHttpError(500, description)}
}

var ErrorSet []error

func ManageMultipleError(err error) {
	if err == nil {
		return
	}
	if ErrorSet == nil {
		ErrorSet = make([]error, 0)
	}
	ErrorSet = append(ErrorSet, err)
}

func ConculdeMultipleError() error {
	if ErrorSet == nil {
		return nil
	}
	var errorStr string
	for _, e := range ErrorSet {
		errorStr += e.Error() + "\n"
	}
	errorStr = errorStr[:len(errorStr)-2]
	ErrorSet = nil
	return NewHttpError(500, errorStr)
}
