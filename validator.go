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

// RegisterRule registers a new rule into the validator.
// It returns an error if the rule name is already registered.
func (v *Validator) RegisterRule(rule Rule) error {
	for _, r := range v.rules {
		if r.Name() == rule.Name() {
			return &DuplicatedRuleError{RuleName: rule.Name()}
		}
	}

	v.rules = append(v.rules, rule)

	return nil
}
