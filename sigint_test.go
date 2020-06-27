package sigint

import (
	"encoding/json"
	"testing"

	"fmt"
)

// Signature ...
type signature struct {
	Signature string // Signature byte sequence.
	Encoding  string // Signature encoding, e.g. Hexadecimal, ASCII, PRONOM.

	// TODO: Do we need relativity?
	Relativity string // Position relative to beginning or end of file, or elsewhere.
	Comment    string
	Fail       bool
}

func TestParse(t *testing.T) {
	var sigs []signature
	err := json.Unmarshal([]byte(testPatterns), &sigs)
	if err != nil {
		t.Error("Failed to load fixtures:", err)
	}
	for _, sig := range sigs {
		_, err := Parse(sig.Signature, LookupEncoding(sig.Encoding))
		if err != nil && sig.Fail != true {
			t.Error("Failed to parse signature:", err)
		}
	}
	fmt.Println("Testing complete...")
}
