package rule

import (
	"testing"
)

func TestMinRuleName(t *testing.T) {
	var r Min

	if r.Name() != "min" {
		t.Error("unexpected rule name")
	}
}

func TestErrMinMessage(t *testing.T) {
	err := &ErrMin{
		FieldName: "age",
		MinValue:  "18",
	}

	expected := "age must be at least 18."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestMinValidate(t *testing.T) {
	var r Min

	// Test for a value greater than the minimum (should return nil).
	if err := r.Validate("age", "25", "18"); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for a value equal to the minimum (should return nil).
	if err := r.Validate("age", "18", "18"); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for a value less than the minimum (should return ErrMin).
	err := r.Validate("age", "15", "18")
	expected := &ErrMin{FieldName: "age", MinValue: "18"}

	if err == nil {
		t.Error("expected ErrMin, got nil")
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
