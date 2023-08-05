package rule

import "testing"

func TestNumberRuleName(t *testing.T) {
	var r Number

	if r.Name() != "number" {
		t.Error("unexpected rule name")
	}
}

func TestErrNumberMessage(t *testing.T) {
	err := &ErrNumber{
		FieldName: "age",
	}

	expected := "age must only contain numbers."

	if err.Error() != expected {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected, err.Error())
	}
}

func TestNumberValidate(t *testing.T) {
	var r Number

	// Test for a non-empty value (should return nil).
	if err := r.Validate("age", "50", ""); err != nil {
		t.Error("unexpected error:", err)
	}

	// Test for an non-number value (should return ErrNumber).
	err := r.Validate("age", "fourty", "")
	expected := &ErrNumber{FieldName: "age"}

	if err == nil {
		t.Error("expected ErrNumber, got nil")
	} else if err.Error() != expected.Error() {
		t.Errorf("Error message mismatch.\nExpected: %s\nGot: %s", expected.Error(), err.Error())
	}
}
