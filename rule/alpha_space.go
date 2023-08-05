package rule

type (
	AlphaSpace struct{}

	ErrAlphaSpace struct {
		FieldName string
	}
)

func (r *AlphaSpace) Name() string {
	return "alpha_space"
}

func (e *ErrAlphaSpace) Error() string {
	return e.FieldName + " must only contain letters and spaces."
}

func (r *AlphaSpace) Validate(field, value, _ string) error {
	if !alphaSpaceRegex.MatchString(value) {
		return &ErrAlphaSpace{FieldName: field}
	}

	return nil
}
