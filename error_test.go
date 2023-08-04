package blu

import (
	"reflect"
	"testing"
)

func TestErrorMessages(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "error message: ErrorDuplicatedRule",
			err:      &ErrorDuplicatedRule{RuleName: "alpha_space"},
			expected: "error: duplicated rule alpha_space. the rule already exists.",
		},
		{
			name:     "error message: ErrorUnexportedField",
			err:      &ErrorUnexportedField{Field: "name"},
			expected: "error: unexported field name encountered.",
		},
		{
			name:     "error message: ErrorInvalidInput",
			err:      &ErrorInvalidInput{reflect.TypeOf(true)},
			expected: "error: invalid input of type bool. input is not a struct.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err.Error() != tc.expected {
				t.Errorf("expected error message: %s, got: %s", tc.expected, tc.err.Error())
			}
		})
	}
}
