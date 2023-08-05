package rule

type (
	Number struct{}

	ErrNumber struct {
		FieldName string
	}
)

func (r *Number) Name() string {
	return "number"
}

func (e *ErrNumber) Error() string {
	return e.FieldName + " must only contain numbers."
}

func (r *Number) Validate(field, value, _ string) error {
	if !numberRegex.MatchString(value) {
		return &ErrNumber{FieldName: field}
	}

	return nil
}
