package rule

import "testing"

func TestNumericRuleName(t *testing.T) {
	var r Numeric

	if r.Name() != "numeric" {
		t.Error("unexpected rule name")
	}
}

func TestErrNumericMessage(t *testing.T) {
	err := &ErrNumeric{
		FieldName: "price",
	}

	expected := "price must only contain numeric values."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestNumericValidate(t *testing.T) {
	var r Numeric

	// Test for a non-empty value (should return nil).
	if err := r.Validate("price", "10.89", ""); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an invalid numeric value (should return ErrNumeric).
	err := r.Validate("price", "10-", "")
	expected := &ErrNumeric{FieldName: "price"}

	if err == nil {
		t.Error("expected ErrNumeric, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
