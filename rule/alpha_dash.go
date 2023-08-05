package rule

type (
	AlphaDash struct{}

	ErrAlphaDash struct {
		FieldName string
	}
)

func (r *AlphaDash) Name() string {
	return "alpha_dash"
}

func (e *ErrAlphaDash) Error() string {
	return e.FieldName + " must only contain letters, numbers, dashes, and underscores."
}

func (r *AlphaDash) Validate(field, value, _ string) error {
	if !alphaDashRegex.MatchString(value) {
		return &ErrAlphaDash{FieldName: field}
	}

	return nil
}
