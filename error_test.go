package blu

import (
	"reflect"
	"testing"
)

func TestErrorMessages(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "error message: DuplicatedRuleError",
			err:      &DuplicatedRuleError{RuleName: "alpha_space"},
			expected: "duplicated rule: rule alpha_space already exist.",
		},
		{
			name:     "error message: UnexportedFieldError",
			err:      &UnexportedFieldError{Field: "name"},
			expected: "unexported field encountered for field: name",
		},
		{
			name:     "error message: InvalidInputError",
			err:      &InvalidInputError{reflect.TypeOf(true)},
			expected: "invalid input of type bool: input is not a struct",
		},
		{
			name:     "error message: ValidationError",
			err:      &ValidationError{},
			expected: "validation error: some fields have validation issues",
		},
		{
			name:     "error message: EmptyFieldValueError",
			err:      &EmptyFieldValueError{FieldName: "full_name"},
			expected: "this field full_name can't empty",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.err.Error() != tc.expected {
				t.Errorf("expected error message: %s, got: %s", tc.expected, tc.err.Error())
			}
		})
	}
}
