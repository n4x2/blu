package blu

// Rule represents a validation rule for validating field values.
type Rule interface {
	// Name return the name of the validation rule.
	Name() string
	// Validate performs validation on the field value with the provided parameters.
	// It returns an error if the validation fails; otherwise, it returns nil.
	Validate(field, value, param string) error
}
