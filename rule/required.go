package rule

type (
	Required struct{}

	ErrRequired struct {
		FieldName string
	}
)

func (r *Required) Name() string {
	return "required"
}

func (e *ErrRequired) Error() string {
	return e.FieldName + " is required."
}

func (r *Required) Validate(field, value, _ string) error {
	if value == "" {
		return &ErrRequired{FieldName: field}
	}

	return nil
}
