package rule

import "regexp"

const (
	paramSeparator = ","
)

const (
	alphaDashPattern = "^[a-zA-Z0-9_-]+$"
	alphaNumPattern  = "^[a-zA-Z0-9]+$"
	numberPattern    = "^[0-9]+$"
	numericPattern   = "^[-+]?\\d+(\\.\\d+)?$"
)

var (
	alphaDashRegex = regexp.MustCompile(alphaDashPattern)
	alphaNumRegex  = regexp.MustCompile(alphaNumPattern)
	numberRegex    = regexp.MustCompile(numberPattern)
	numericRegex   = regexp.MustCompile(numericPattern)
)
