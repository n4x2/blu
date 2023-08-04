package blu

import (
	"fmt"
	"reflect"
	"strings"
)

// serialize serializes struct into a list of fields.
func serialize(s interface{}) []Field {
	var fields []Field

	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		var field Field

		field.Name = parseName(t.Field(i))
		field.Value = parseValue(v.Field(i))
		field.Tags = parseTags(t.Field(i).Tag.Get(defaultTagName))

		fields = append(fields, field)
	}

	return fields
}

// parseName parse field name or JSON tag name if exists.
func parseName(f reflect.StructField) string {
	if name, exist := f.Tag.Lookup(jsonTagName); exist {
		return name
	}

	return f.Name
}

// parseTags parse tag name and tag parameter if exists.
func parseTags(t string) []Tag {
	var tags []Tag

	tagList := strings.Split(t, separatorTag)
	for _, tag := range tagList {
		var tagParsed Tag

		parts := strings.Split(tag, pairSeparatorTag)

		tagParsed.Name = parts[indexTagName]
		if len(parts) > 1 {
			tagParsed.Param = parts[indexTagParams]
		}

		tags = append(tags, tagParsed)
	}

	return tags
}

// parseValue parse field value.
// If field is non-string and empty, it return empty string.
func parseValue(f reflect.Value) string {
	if f.IsZero() {
		return ""
	}

	return fmt.Sprintf("%v", f.Interface())
}
