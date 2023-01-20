// internal/sanitizer.go

package internal

import (
	"log"
	"net/url"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Returns a string without accents
func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		log.Fatal(e)
	}
	return output
}

// Returns an all lower-case string in the form of `a-lower-case-string`
func GetSanitizedCommonName(s string) string {
	clean := RemoveAccents(s)
	str := url.PathEscape(strings.ToLower(clean))
	re := regexp.MustCompile("%[0-9A-Fa-f]{2}")
	return re.ReplaceAllString(str, "-")
}
