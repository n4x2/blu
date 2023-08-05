package rule

import "regexp"

const (
	paramSeparator = ","
)

const (
	numberPattern = "^[0-9]+$"
)

var (
	numberRegex = regexp.MustCompile(numberPattern)
)
