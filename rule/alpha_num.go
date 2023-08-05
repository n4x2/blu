package rule

type (
	AlphaNum struct{}

	ErrAlphaNum struct {
		FieldName string
	}
)

func (r *AlphaNum) Name() string {
	return "alpha_num"
}

func (e *ErrAlphaNum) Error() string {
	return e.FieldName + " must only contain letters and numbers."
}

func (r *AlphaNum) Validate(field, value, _ string) error {
	if !alphaNumRegex.MatchString(value) {
		return &ErrAlphaNum{FieldName: field}
	}

	return nil
}
