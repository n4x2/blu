package rule

import "testing"

func TestAlphaSpaceRuleName(t *testing.T) {
	var r AlphaSpace

	if r.Name() != "alpha_space" {
		t.Error("unexpected rule name")
	}
}

func TestErrAlphaSpaceMessage(t *testing.T) {
	err := &ErrAlphaSpace{
		FieldName: "example",
	}

	expected := "example must only contain letters and spaces."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestAlphaSpaceValidate(t *testing.T) {
	var r AlphaSpace

	// Test for a non-empty value (should return nil).
	if err := r.Validate("full_name", "john doe", ""); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an value with character (should return ErrAlphaSpace).
	err := r.Validate("full_name", "prof. john doe", "")
	expected := &ErrAlphaSpace{FieldName: "full_name"}

	if err == nil {
		t.Error("expected ErrAlphaSpace, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
