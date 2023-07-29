package blu

import (
	"fmt"
	"reflect"
)

type (
	// DuplicatedRuleError an error type for attempting register new rule.
	DuplicatedRuleError struct {
		RuleName string
	}

	// InvalidaInputError an error type indicating an issue with input validation.
	InvalidInputError struct {
		Type reflect.Type
	}

	// UnexportedFieldError an error type for attempting to validate an unexported field.
	UnexportedFieldError struct {
		Field string
	}

	// ValidationError storing validation errors for fields as key-value pairs.
	ValidationError map[string][]string
)

// Error returns an error message indicating that a rule with the same name already exists.
func (e *DuplicatedRuleError) Error() string {
	return "duplicated rule: rule " + e.RuleName + " already exist."
}

// Error return for invalid input (non-struct).
func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input of type %s: input is not a struct", e.Type)
}

// Error returns an error message indicating unexported field (lowercase).
func (e *UnexportedFieldError) Error() string {
	return "unexported field encountered for field: " + e.Field
}

// Error returns a error message indicating that there are validation issues.
func (e *ValidationError) Error() string {
	return "validation error: some fields have validation issues"
}
