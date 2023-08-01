package blu

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Validator contains the validation rules.
type Validator struct {
	rules []Rule // List of registered rules
}

// NewValidator returns a new Validator instance.
func NewValidator() *Validator {
	return &Validator{
		rules: []Rule{},
	}
}

// RegisterRule registers a new rule into the validator.
// It returns an error if the rule name is already registered.
func (v *Validator) RegisterRule(rule Rule) error {
	for _, r := range v.rules {
		if r.Name() == rule.Name() {
			return &DuplicatedRuleError{RuleName: rule.Name()}
		}
	}

	v.rules = append(v.rules, rule)

	return nil
}

// Serialize serializes value of struct into a list of fields with their tags.
// It returns an error if struct contain unexported field (lowercase).
func (v *Validator) Serialize(value reflect.Value) ([]Field, error) {
	var fields []Field

	for i := 0; i < value.NumField(); i++ {
		var field Field

		// Get the field name.
		field.Name = value.Type().Field(i).Name

		// Get JSON name if it exists.
		if jsonName, exists := value.Type().Field(i).Tag.Lookup(defaultJSONTag); exists {
			field.Name = strings.ToLower(jsonName)
		}

		// Check if field is unable to export (lowercase).
		if value.Type().Field(i).PkgPath != emptyString {
			return nil, &UnexportedFieldError{Field: field.Name}
		}

		// Get the field value and convert it to a string.
		field.Value = fmt.Sprintf("%v", value.Field(i).Interface())

		// Check if the non-string value is zero then set value to empty string.
		if value.Field(i).IsZero() {
			field.Value = emptyString
		}

		var tagsParsed []Tag

		// Parse tags if exists.
		tags, exists := value.Type().Field(i).Tag.Lookup(defaultTag)
		if exists && tags != emptyString {
			tagList := strings.Split(tags, defaultTagSeparator)

			for _, tag := range tagList {
				var tagParsed Tag

				tagParsed.Name = tag
				tagParsed.Param = emptyString

				// Parsing tag parameter if exists.
				if strings.Contains(tag, defaultTagPairSeparator) {
					pair := strings.SplitN(tag, defaultTagPairSeparator, maxTagSplit)

					tagParsed.Name = pair[defaultTagNameIndex]
					tagParsed.Param = pair[defaultTagParamIndex]
				}

				tagsParsed = append(tagsParsed, tagParsed)
			}
		}

		field.Tags = tagsParsed

		fields = append(fields, field)
	}

	return fields, nil
}

// Validate validates the field value of the given struct based on tags.
// It return an error if the input is not a struct.
// If there are validation errors for the struct fields,
// it return a joined error that contains all the validation errors.
// It return nil if validation success.
func (v *Validator) Validate(s interface{}) error {
	value := reflect.ValueOf(s)

	if value.Kind() != reflect.Struct {
		return &InvalidInputError{Type: reflect.TypeOf(s)}
	}

	// Serialize struct.
	fields, err := v.Serialize(value)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	var validationErrors error

	// Validating fields.
	for _, field := range fields {
		for _, tag := range field.Tags {
			// Skip validation if contain "optional" tag, except value not empty.
			if tag.Name == defaultOptionalTag {
				if field.Value == emptyString {
					break
				}
			}

			for _, rule := range v.rules {
				if rule.Name() == tag.Name {
					err := rule.Validate(field.Name, field.Value, tag.Param)
					if err != nil {
						validationErrors = errors.Join(validationErrors, err)
					}
				}
			}
		}
	}

	return validationErrors
}
