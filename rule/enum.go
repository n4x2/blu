package rule

import "strings"

type (
	Enum struct{}

	ErrEnum struct {
		FieldValue string
	}
)

func (r *Enum) Name() string {
	return "enum"
}

func (e *ErrEnum) Error() string {
	return "selected " + e.FieldValue + " is invalid."
}

func (r *Enum) Validate(_, value, param string) error {
	options := strings.Split(param, paramSeparator)
	for _, option := range options {
		if value == option {
			return nil
		}
	}

	return &ErrEnum{FieldValue: value}
}
