package rule

import (
	"errors"
	"strconv"
)

type (
	Max struct{}

	ErrMax struct {
		FieldName string
		MaxValue  string
	}
)

func (r *Max) Name() string {
	return "max"
}

func (e *ErrMax) Error() string {
	return e.FieldName + " must not be greater than " + e.MaxValue + "."
}

func (r *Max) Validate(field, value, param string) error {
	minValue, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return errors.Unwrap(err)
	}

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.Unwrap(err)
	}

	if v > minValue {
		return &ErrMax{FieldName: field, MaxValue: param}
	}

	return nil
}
