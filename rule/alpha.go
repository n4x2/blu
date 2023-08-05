package rule

type (
	Alpha struct{}

	ErrAlpha struct {
		FieldName string
	}
)

func (r *Alpha) Name() string {
	return "alpha"
}

func (e *ErrAlpha) Error() string {
	return e.FieldName + " must only contain letters."
}

func (r *Alpha) Validate(field, value, _ string) error {
	if !alphaRegex.MatchString(value) {
		return &ErrAlpha{FieldName: field}
	}

	return nil
}
