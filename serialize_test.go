package blu

import (
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name: "struct with json name",
			input: struct {
				Name    string `json:"name" validate:"required|alpha_space"`
				Address string `json:"address" validate:"required_if:name,valid|alpha_dash"`
			}{
				Name:    "John Doe",
				Address: "Jl. Palagan KM. 5",
			},
			expected: []Field{
				{
					Name:  "name",
					Value: "John Doe",
					Tags: []Tag{
						{Name: "required", Param: ""},
						{Name: "alpha_space", Param: ""},
					},
				},
				{
					Name:  "address",
					Value: "Jl. Palagan KM. 5",
					Tags: []Tag{
						{Name: "required_if", Param: "name,valid"},
						{Name: "alpha_dash", Param: ""},
					},
				},
			},
		},
		{
			name: "struct without json name",
			input: struct {
				Name string `validate:"required|alpha_space"`
				Age  int
			}{
				Name: "John Doe",
			},
			expected: []Field{
				{
					Name:  "Name",
					Value: "John Doe",
					Tags: []Tag{
						{Name: "required", Param: ""},
						{Name: "alpha_space", Param: ""},
					},
				},
				{
					Name: "Age",
					Tags: []Tag{{}},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := serialize(tc.input)
			if !reflect.DeepEqual(tc.expected, f) {
				t.Errorf("error want %v got %v", tc.expected, f)
			}
		})
	}
}
