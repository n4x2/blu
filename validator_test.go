package blu

import (
	"errors"
	"testing"
)

var ErrExampleRule = errors.New("value must be example.")

type (
	ExampleRule struct{}
)

func (r *ExampleRule) Name() string {
	return "example"
}

func (r *ExampleRule) Validate(_, value, _ string) error {
	if value != r.Name() {
		return ErrExampleRule
	}

	return nil
}

func TestRegisterRule(t *testing.T) {
	v := NewValidator()

	err := v.RegisterRule(&ExampleRule{})
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	// Attempt to register the same rule again.
	err = v.RegisterRule(&ExampleRule{})
	if err == nil {
		t.Error("expected ErrorDuplicatedRule, but got nil")
	} else {
		var dupErr *ErrorDuplicatedRule
		if !errors.As(err, &dupErr) {
			t.Errorf("unexpected error type: %T, expected *ErrorDuplicatedRule", err)
		}
	}
}
