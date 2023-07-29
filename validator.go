package blu

// Validator contains the validation rules.
type Validator struct {
	rules []Rule // List of registered rules
}

// NewValidator returns a new Validator instance.
func NewValidator() *Validator {
	return &Validator{
		rules: []Rule{},
	}
}
