package rule

import "regexp"

const (
	paramSeparator = ","
)

const (
	alphaDashPattern = "^[a-zA-Z0-9_-]+$"
	numberPattern    = "^[0-9]+$"
	numericPattern   = "^[-+]?\\d+(\\.\\d+)?$"
)

var (
	alphaDashRegex = regexp.MustCompile(alphaDashPattern)
	numberRegex    = regexp.MustCompile(numberPattern)
	numericRegex   = regexp.MustCompile(numericPattern)
)
