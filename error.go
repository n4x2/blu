package blu

import (
	"fmt"
	"reflect"
)

type (
	// ErrorDuplicatedRule an error type for attempting register new rule.
	ErrorDuplicatedRule struct {
		RuleName string
	}

	// InvalidaInputError an error type indicating an issue with input validation.
	ErrorInvalidInput struct {
		Type reflect.Type
	}

	// ErrorUnexportedField an error type for attempting to validate an unexported field.
	ErrorUnexportedField struct {
		Field string
	}
)

// Error returns an error message indicating that a rule with the same name already exists.
func (e *ErrorDuplicatedRule) Error() string {
	return "error: duplicated rule " + e.RuleName + ". the rule already exists."
}

// Error return for invalid input (non-struct).
func (e *ErrorInvalidInput) Error() string {
	return fmt.Sprintf("error: invalid input of type %s. input is not a struct.", e.Type)
}

// Error returns an error message indicating unexported field (lowercase).
func (e *ErrorUnexportedField) Error() string {
	return "error: unexported field " + e.Field + " encountered."
}
