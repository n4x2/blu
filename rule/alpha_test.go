package rule

import "testing"

func TestAlphaRuleName(t *testing.T) {
	var r Alpha

	if r.Name() != "alpha" {
		t.Error("unexpected rule name")
	}
}

func TestErrAlphaMessage(t *testing.T) {
	err := &ErrAlpha{
		FieldName: "example",
	}

	expected := "example must only contain letters."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestAlphaValidate(t *testing.T) {
	var r Alpha

	// Test for a non-empty value (should return nil).
	if err := r.Validate("first_name", "john", ""); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an value with number (should return ErrAlpha).
	err := r.Validate("first_name", "john123", "")
	expected := &ErrAlpha{FieldName: "first_name"}

	if err == nil {
		t.Error("expected ErrAlpha, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
