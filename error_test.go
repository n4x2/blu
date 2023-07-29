package blu

import (
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
