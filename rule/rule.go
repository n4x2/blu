package rule

import "regexp"

const (
	paramSeparator = ","
)

const (
	alphaDashPattern  = "^[a-zA-Z0-9_-]+$"
	alphaNumPattern   = "^[a-zA-Z0-9]+$"
	alphaSpacePattern = "^[a-zA-Z ]+$"
	alphaPattern      = "^[a-zA-Z]+$"
	numberPattern     = "^[0-9]+$"
	numericPattern    = "^[-+]?\\d+(\\.\\d+)?$"
)

var (
	alphaDashRegex  = regexp.MustCompile(alphaDashPattern)
	alphaNumRegex   = regexp.MustCompile(alphaNumPattern)
	alphaSpaceRegex = regexp.MustCompile(alphaSpacePattern)
	alphaRegex      = regexp.MustCompile(alphaPattern)
	numberRegex     = regexp.MustCompile(numberPattern)
	numericRegex    = regexp.MustCompile(numericPattern)
)
