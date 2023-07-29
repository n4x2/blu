// blu provide struct value validation based on tags
package blu

const (
	defaultJSONTag          = "json"     // Default name for JSON tag.
	defaultTag              = "validate" // Default name for validation tag.
	defaultTagNameIndex     = 0          // Index of tag name in a tag string (e.g., "validate:name=param").
	defaultTagPairSeparator = "="        // Separator between tag name and parameter.
	defaultTagParamIndex    = 1          // Index of tag parameter in a tag string (e.g., "validate:name=param").
	defaultTagSeparator     = ","        // Separator between multiple tags in a field.
	emptyString             = ""         // Empty string constant.
	maxTagSplit             = 2          // Maximum number of splits when parsing a tag.
)
