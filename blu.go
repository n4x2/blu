// blu is struct value validation based on tags.
package blu

const (
	emptyTag    = "empty" // Tag name used to specify a field that should be considered empty.
	emptyString = ""      // An empty string used for comparison in some validation rules.
)

const (
	defaultTagName   = "validate" // Default tag name used for validation if no specific tag is provided.
	jsonTagName      = "json"     // Tag name used to indicate JSON field names in the struct.
	indexTagName     = 0          // Index of the tag name in the tag string when split by pairSeparatorTag.
	indexTagParams   = 1          // Index of the tag parameters in the tag string when split by pairSeparatorTag.
	separatorTag     = "|"        // Separator used to split multiple tags in a single struct tag.
	pairSeparatorTag = ":"        // Separator used to split the tag name and its parameters.
)

// supportedRules default supported rules.
var supportedRules = []Rule{}
