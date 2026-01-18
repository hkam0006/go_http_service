package products

import (
	"encoding/json"
	"io"
)

type Validator interface {
	Validate(io.ReadCloser, any) error
}

type validator struct {}

func (v *validator) Validate(body io.ReadCloser, format any) error {
	if err := json.NewDecoder(body).Decode(format); err != nil {
		return err
	}

	return nil
}

func NewValidator() Validator {
	return &validator{}
}
