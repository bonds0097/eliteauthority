package webUtils

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// Generates a slug from an ASCII-printable sting.
func GenerateSlug(in string) (out string, err error) {
	validChars, _ := regexp.Compile(`[\w_\-~\.]`)
	whitespace, _ := regexp.Compile(`\s+`)
	in = whitespace.ReplaceAllLiteralString(in, "-")
	outArray := validChars.FindAllString(in, -1)
	out = strings.ToLower(strings.Join(outArray, ""))

	if len(out) == 0 {
		err = errors.New("Failed to generate slug, no valid characters in input.")
	}

	return
}

// Remove leading, trailing whitespace and replace multiple spaces with a single space.
func NormalizeWhitespace(in string) (out string) {
	whitespace, _ := regexp.Compile(`\s{2,}`)
	out = strings.TrimSpace(whitespace.ReplaceAllString(in, " "))

	return
}

func IsAsciiPrintable(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
