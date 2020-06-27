package sigint

import (
	"strings"
)

const (
	// UnknownEncoding provides us with a default to work with.
	UnknownEncoding = iota
	// HexEncoding describes magic numbers written in plain-hexadecimal.
	HexEncoding
	// PronomEncoding describe PRONOM based file format signatures.
	PronomEncoding
	// PerlEncoding describe PERL regular expression encoded signatures.
	PerlEncoding
	// ASCIIEncoding encoded patterns are those written entirely in plain ASCII.
	ASCIIEncoding
)

// LookupEncoding will return a best-guess encoding type for a supplied
// encoding string.
//
// TODO: I really don't know if we need this... but it's a convenience...
//
func LookupEncoding(encoding string) int {
	encoding = strings.ToLower(encoding)
	switch encoding {
	case "hexadecimal":
		return HexEncoding
	case "pronom internal signature":
		return PronomEncoding
	case "perl compatible regular expressions":
		return PerlEncoding
	case "ascii":
		return ASCIIEncoding
	}
	return UnknownEncoding
}
