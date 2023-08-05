package rule

import "regexp"

const (
	paramSeparator = ","
)

const (
	numberPattern  = "^[0-9]+$"
	numericPattern = "^[-+]?\\d+(\\.\\d+)?$"
)

var (
	numberRegex  = regexp.MustCompile(numberPattern)
	numericRegex = regexp.MustCompile(numericPattern)
)
