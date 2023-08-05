package rule

import "testing"

func TestAlphaNumRuleName(t *testing.T) {
	var r AlphaNum

	if r.Name() != "alpha_num" {
		t.Error("unexpected rule name")
	}
}

func TestErrAlphaNumMessage(t *testing.T) {
	err := &ErrAlphaNum{
		FieldName: "example",
	}

	expected := "example must only contain letters and numbers."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestAlphaNumValidate(t *testing.T) {
	var r AlphaNum

	// Test for a non-empty value (should return nil).
	if err := r.Validate("coupon_code", "rR902mk", ""); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an value with special character (should return ErrAlphaNum).
	err := r.Validate("coupon_code", "rR902mk!%", "")
	expected := &ErrAlphaNum{FieldName: "coupon_code"}

	if err == nil {
		t.Error("expected ErrAlphaNum, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
