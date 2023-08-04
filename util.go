package blu

import "reflect"

// isValidInput checks if the provided value is a valid input for validation.
// It returns true if the value is a struct, otherwise false.
func isValidInput(s interface{}) bool {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		return false
	}

	return true
}

// isExportableStruct checks if the provided struct is exportable (all fields are exported).
// It returns true and an empty string if the struct is exportable.
// If the struct has unexported fields (fields starting with lowercase letters),
// it returns false and the name of the first unexported field encountered.
func isExportableStruct(s interface{}) (bool, string) {
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		// Check if field is unable to export (lowercase).
		if v.Type().Field(i).PkgPath != "" {
			return false, v.Type().Field(i).Name
		}
	}

	return true, ""
}
