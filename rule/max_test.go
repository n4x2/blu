package rule

import (
	"testing"
)

func TestMaxRuleName(t *testing.T) {
	var r Max

	if r.Name() != "max" {
		t.Error("unexpected rule name")
	}
}

func TestErrMaxMessage(t *testing.T) {
	err := &ErrMax{
		FieldName: "age",
		MaxValue:  "18",
	}

	expected := "age must not be greater than 18."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestMaxValidate(t *testing.T) {
	var r Max

	// Test for a value less than the maximum (should return nil).
	if err := r.Validate("age", "50", "65"); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for a value equal to the maximum (should return nil).
	if err := r.Validate("age", "65", "65"); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for a value greater than the maximum (should return ErrMax).
	err := r.Validate("age", "70", "65")
	expected := &ErrMax{FieldName: "age", MaxValue: "65"}

	if err == nil {
		t.Error("expected ErrMax, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}

	// Test for invalid value.
	err = r.Validate("age", "18 years old", "18")
	if err == nil {
		t.Error("expected error but got nil")
	}

	// Test for invalid parameter.
	err = r.Validate("age", "18", "true")
	if err == nil {
		t.Error("expected error but got nil")
	}
}
