package blu

import "testing"

type (
	Exportable struct {
		Name string
	}

	Unexportable struct {
		name string
	}
)

func TestValidInput(t *testing.T) {
	var s Exportable

	ok := isValidInput(s)
	if !ok {
		t.Errorf("expected true but got %v", ok)
	}
}

func TestInvalidInput(t *testing.T) {
	ok := isValidInput(true)
	if ok {
		t.Errorf("expected false but got %v", ok)
	}
}

func TestExportableStruct(t *testing.T) {
	var s Exportable

	ok, _ := isExportableStruct(s)
	if !ok {
		t.Errorf("expected no error")
	}
}

func TestUnexportableStruct(t *testing.T) {
	var s Unexportable

	ok, _ := isExportableStruct(s)
	if ok {
		t.Errorf("expected error")
	}
}
