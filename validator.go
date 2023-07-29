package blu

import (
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
