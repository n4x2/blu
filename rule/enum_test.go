package rule

import "testing"

func TestEnumRuleName(t *testing.T) {
	var r Enum

	if r.Name() != "enum" {
		t.Error("unexpected rule name")
	}
}

func TestErrEnumMessage(t *testing.T) {
	err := &ErrEnum{
		FieldValue: "example",
	}

	expected := "selected example is invalid."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestEnumValidate(t *testing.T) {
	var r Enum

	// Test for a non-empty value (should return nil).
	if err := r.Validate("", "one", "one,two,three"); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an invalid value (should return ErrEnum).
	err := r.Validate("", "ten", "one,two,three")
	expected := &ErrEnum{FieldValue: "ten"}

	if err == nil {
		t.Error("expected ErrEnum, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
