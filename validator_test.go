package blu

import "testing"

type CustomRule struct {
	NameVal string
}

func (r *CustomRule) Name() string {
	return r.NameVal
}

func (r *CustomRule) Validate(field, value, params string) error {
	return nil
}

func TestRegisterRule(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name         string
		initialRules []Rule
		newRule      Rule
		expectedErr  error
	}

	testCases := []testCase{
		{
			name:         "register new rule",
			initialRules: nil,
			newRule:      &CustomRule{NameVal: "rule"},
			expectedErr:  nil,
		},
		{
			name: "register duplicated rule",
			initialRules: []Rule{
				&CustomRule{NameVal: "rule"},
			},
			newRule:     &CustomRule{NameVal: "rule"},
			expectedErr: &DuplicatedRuleError{RuleName: "rule"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			validator := NewValidator()

			for _, rule := range tc.initialRules {
				validator.RegisterRule(rule)
			}

			err := validator.RegisterRule(tc.newRule)
			if (err == nil && tc.expectedErr != nil) || (err != nil && err.Error() != tc.expectedErr.Error()) {
				t.Errorf("expected error: %v, got error: %v", tc.expectedErr, err)
			}
		})
	}
}
