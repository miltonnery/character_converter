package validator

import (
	"miltonnery/character_converter/dto"
	errorHandling "miltonnery/character_converter/error"
	"strings"
)

type RequestValidator struct {
	request *dto.Request
}

func NewRequestValidator(request *dto.Request) *RequestValidator {
	return &RequestValidator{request: request}
}

func (rv *RequestValidator) OK() error {

	if strings.EqualFold(rv.request.TextoATraducir, "") {
		err := errorHandling.NewInternalErrorWithCustomizedMessage(errorHandling.FieldMissing, "The input text is empty")
		return err
	}

	return nil
}
