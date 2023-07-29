package blu

type (
	// DuplicatedRuleError an error type for attempting register new rule.
	DuplicatedRuleError struct {
		RuleName string
	}
)

// Error returns an error message indicating that a rule with the same name already exists.
func (e *DuplicatedRuleError) Error() string {
	return "duplicated rule: rule " + e.RuleName + " already exist."
}
