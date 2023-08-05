package rule

import (
	"errors"
	"strconv"
)

type (
	Min struct{}

	ErrMin struct {
		FieldName string
		MinValue  string
	}
)

func (r *Min) Name() string {
	return "min"
}

func (e *ErrMin) Error() string {
	return e.FieldName + " must be at least " + e.MinValue + "."
}

func (r *Min) Validate(field, value, param string) error {
	minValue, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return errors.Unwrap(err)
	}

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.Unwrap(err)
	}

	if v < minValue {
		return &ErrMin{FieldName: field, MinValue: param}
	}

	return nil
}
