package analyzers

import (
	"regexp"
)

// ExtractLibsFromFile selfexplained
func ExtractLibsFromFile(content string, extractor *regexp.Regexp) (findings [][]string, err error) {
	findings = extractor.FindAllStringSubmatch(content, -1)

	return
}
