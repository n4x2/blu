package rule

import "testing"

func TestAlphaDashRuleName(t *testing.T) {
	var r AlphaDash

	if r.Name() != "alpha_dash" {
		t.Error("unexpected rule name")
	}
}

func TestErrAlphaDashMessage(t *testing.T) {
	err := &ErrAlphaDash{
		FieldName: "example",
	}

	expected := "example must only contain letters, numbers, dashes, and underscores."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestAlphaDashValidate(t *testing.T) {
	var r AlphaDash

	// Test for a non-empty value (should return nil).
	if err := r.Validate("username", "john9_r9-19", ""); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an value with space (should return ErrAlphaDash).
	err := r.Validate("username", "john doe", "")
	expected := &ErrAlphaDash{FieldName: "username"}

	if err == nil {
		t.Error("expected ErrAlphaDash, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
