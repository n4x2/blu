package blu

type (
	// DuplicatedRuleError an error type for attempting register new rule.
	DuplicatedRuleError struct {
		RuleName string
	}

	// UnexportedFieldError an error type for attempting to validate an unexported field.
	UnexportedFieldError struct {
		Field string
	}
)

// Error returns an error message indicating that a rule with the same name already exists.
func (e *DuplicatedRuleError) Error() string {
	return "duplicated rule: rule " + e.RuleName + " already exist."
}

// Error returns an error message indicating unexported field (lowercase).
func (e *UnexportedFieldError) Error() string {
	return "unexported field encountered for field: " + e.Field
}
