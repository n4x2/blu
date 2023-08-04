package blu

import (
	"errors"
	"reflect"
)

// Validator contains the validation rules.
type Validator struct {
	rules []Rule // List of registered rules
}

// NewValidator returns a new Validator instance.
func NewValidator() *Validator {
	return &Validator{
		rules: supportedRules,
	}
}

// RegisterRule registers a new rule into the validator.
// It returns an error if the rule name is already registered.
func (v *Validator) RegisterRule(rule Rule) error {
	for _, r := range v.rules {
		if r.Name() == rule.Name() {
			return &ErrorDuplicatedRule{RuleName: rule.Name()}
		}
	}

	v.rules = append(v.rules, rule)

	return nil
}

// ValidateField validates the field value based on the provided tags.
// It returns a joined error that contains all validation errors of the field.
func (v *Validator) ValidateField(f Field) error {
	var errs error

	for _, t := range f.Tags {
		if t.Name == emptyTag && f.Value == emptyString {
			break
		}

		for _, r := range v.rules {
			if r.Name() == t.Name {
				err := r.Validate(f.Name, f.Value, t.Param)
				if err != nil {
					errs = errors.Join(errs, err)
				}
			}
		}
	}

	return errs
}

// Validate validates the field value of the given struct based on tags.
// It returns an error if the input is not a struct.
// It returns an error if struct not exportable.
// If there are validation errors for the struct fields,
// it return a joined error that contains all the validation errors.
// It return nil if validation success.
func (v *Validator) Validate(s interface{}) error {
	if ok := isValidInput(s); !ok {
		return &ErrorInvalidInput{Type: reflect.TypeOf(s)}
	}

	if ok, fieldName := isExportableStruct(s); !ok {
		return &ErrorUnexportedField{Field: fieldName}
	}

	fields := serialize(s)

	var errs error

	for _, field := range fields {
		err := v.ValidateField(field)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}
