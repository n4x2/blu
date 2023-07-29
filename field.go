package blu

type (
	// Field represents struct field with its name, value, and associated tags.
	Field struct {
		Name  string
		Value string
		Tags  []Tag
	}

	// Tag represents a tag with name and parameter.
	Tag struct {
		Name  string
		Param string
	}
)
