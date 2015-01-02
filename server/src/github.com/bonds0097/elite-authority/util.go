package main

import (
	"github.com/bonds0097/webUtils"
	"strings"
)

// Normalizes strings so that they're stored the way we want them.
func normalizeString(in string) (out string) {
	out = strings.Title(webUtils.NormalizeWhitespace(in))

	return
}
