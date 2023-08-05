package rule

type (
	Numeric struct{}

	ErrNumeric struct {
		FieldName string
	}
)

func (r *Numeric) Name() string {
	return "numeric"
}

func (e *ErrNumeric) Error() string {
	return e.FieldName + " must only contain numeric values."
}

func (r *Numeric) Validate(field, value, _ string) error {
	if !numericRegex.MatchString(value) {
		return &ErrNumeric{FieldName: field}
	}

	return nil
}
