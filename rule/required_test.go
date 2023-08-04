package rule

import "testing"

func TestRequiredRuleName(t *testing.T) {
	var r Required

	if r.Name() != "required" {
		t.Error("unexpected rule name")
	}
}

func TestErrRequiredMessage(t *testing.T) {
	err := &ErrRequired{
		FieldName: "username",
	}

	expected := "username is required."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestRequiredValidate(t *testing.T) {
	var r Required

	// Test for a non-empty value (should return nil).
	if err := r.Validate("field", "value", ""); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an empty value (should return ErrRequired).
	err := r.Validate("field", "", "")
	expected := &ErrRequired{FieldName: "field"}

	if err == nil {
		t.Error("expected ErrRequired, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
