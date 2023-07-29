package blu

import (
	"reflect"
	"testing"
)

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
				err := validator.RegisterRule(rule)
				if err != nil {
					t.Fatalf("unexpected error while registering rule: %v", err)
				}
			}

			err := validator.RegisterRule(tc.newRule)
			if (err == nil && tc.expectedErr != nil) || (err != nil && err.Error() != tc.expectedErr.Error()) {
				t.Errorf("expected error: %v, got error: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestSerializeStruct(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    interface{}
		expected []Field
		valid    bool
	}{
		{
			name: "invalid input: unexported field",
			input: struct {
				name string
				age  int
			}{
				name: "John Doe",
				age:  66,
			},
			expected: nil,
			valid:    false,
		},
		{
			name: "valid input: without any tag",
			input: struct {
				Name string
				Age  int
			}{
				Name: "John Doe",
				Age:  66,
			},
			expected: []Field{
				{
					Name:  "Name",
					Value: "John Doe",
				},
				{
					Name:  "Age",
					Value: "66",
				},
			},
			valid: true,
		},
		{
			name: "valid input: with default tag (validate) but empty",
			input: struct {
				Name string `validate:""`
				Age  int    `validate:""`
			}{
				Name: "John Doe",
				Age:  66,
			},
			expected: []Field{
				{
					Name:  "Name",
					Value: "John Doe",
				},
				{
					Name:  "Age",
					Value: "66",
				},
			},
			valid: true,
		},
		{
			name: "valid input: with json and tag",
			input: struct {
				Name string `json:"name" validate:"required,min_length=2,max_length=102"`
				Age  int    `json:"age" validate:"required,max=67"`
			}{
				Name: "John Doe",
				Age:  66,
			},
			expected: []Field{
				{
					Name:  "name",
					Value: "John Doe",
					Tags: []Tag{
						{Name: "required"},
						{Name: "min_length", Param: "2"},
						{Name: "max_length", Param: "102"},
					},
				},
				{
					Name:  "age",
					Value: "66",
					Tags: []Tag{
						{Name: "required"},
						{Name: "max", Param: "67"},
					},
				},
			},
			valid: true,
		},
		{
			name: "valid input: non string empty",
			input: struct {
				Name string `json:"name" validate:"required,min_length=2,max_length=102"`
				Age  int    `json:"age" validate:"required,max=67"`
			}{
				Name: "John Doe",
			},
			expected: []Field{
				{
					Name:  "name",
					Value: "John Doe",
					Tags: []Tag{
						{Name: "required"},
						{Name: "min_length", Param: "2"},
						{Name: "max_length", Param: "102"},
					},
				},
				{
					Name:  "age",
					Value: "",
					Tags: []Tag{
						{Name: "required"},
						{Name: "max", Param: "67"},
					},
				},
			},
			valid: true,
		},
	}

	v := NewValidator()

	for _, tc := range testCases {
		tc := tc // Create a local copy of the loop variable.
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			val := reflect.ValueOf(tc.input)
			fields, err := v.Serialize(val)

			if tc.valid {
				if err != nil {
					t.Errorf("expected no error for valid input, but got: %v", err)
				}
			} else {
				if err == nil {
					t.Error("expected error for invalid input, but got nil")
				}
			}

			if !reflect.DeepEqual(fields, tc.expected) {
				t.Errorf("unexpected result.\nexpected: %+v\nGot: %+v", tc.expected, fields)
			}
		})
	}
}
