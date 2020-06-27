package sigint

import (
	"encoding/hex"
	"strings"

	"fmt"
)

// ParseNoEncoding returns something sensible from a string that has no
// encoding...
func ParseNoEncoding(sig string) {
	fmt.Println("here")
}

// Parse will take a signature and do something useful with it...
func Parse(signature string, encoding int) (string, error) {
	switch encoding {
	case HexEncoding:
		err := HexParse(signature)
		if err != nil {
			// Try to guess the signature...
			fmt.Println("Error", err)
		}
	case ASCIIEncoding:
		_ = ASCIIParser(signature)
	}

	return "", nil
}

// preprocessHex will perform some basic operations on a hexadecimal string to
// give it a greater chance of being decoded.
func preprocessHex(signature string) string {
	// PRE-1: Remove spaces from HEX e.g. `A1 B1` -> `A1B1`.
	return strings.Replace(signature, " ", "", -1)
}

// HexParse will take a hexadecimal based signature and do something useful
// with it...
func HexParse(signature string) error {
	signature = preprocessHex(signature)
	// Validate the hexadecimal
	_, err := hex.DecodeString(signature)
	return err
}

// ASCIIParser returns a hexadecimal representation of a signature written
// using ASCII encoding.
func ASCIIParser(signature string) string {
	return hex.EncodeToString([]byte(signature))
}

/*
	for k, wd := range mappings.WikidataMapping {
		if len(wd.Signatures) > 1 {
			// WIKIDATA TODO: We remove additional signatures we're not yet
			// capable of handling with the Wikidata work. We will add this
			// back in for production release.
			mappings.WikidataMapping[k] = mappings.DeleteSignatures(&wd)
		}
		if len(wd.Signatures) > 0 {
			for signature := range wd.Signatures {
				_, err := isHex(wd.Signatures[signature].Signature)
				if err != nil {
					// WIKIDATA TODO: Any other signatures which are not
					// well-formed will be removed from the collection for now
					// as well and will be added back in for production
					// release.
					mappings.WikidataMapping[k] = mappings.DeleteSignatures(&wd)
				} else {
					// Optionally create Skeleton files from the data...
					roySkeletons := false
					if roySkeletons {
						hexRepr, _ := hex.DecodeString(wd.Signatures[signature].Signature)
						extension := wd.Extension[0]
						fname := ""
						if extension == "" {
							fname = fmt.Sprintf("%s", wd.ID)
						} else {
							fname = fmt.Sprintf("%s.%s", wd.ID, extension)
						}
						fname = filepath.Join("wikidata-skeleton-suite", fname)
						f, err := os.Create(fname)
						if err == nil {
							defer f.Close()
							f.Write([]byte(hexRepr))
						} else {
							fmt.Println("ERROR OPENING FILE", err)
						}

					}
				}
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Roy (Wikidata): Condensed SPARQL results: %d\n", len(mappings.WikidataMapping))
	fmt.Fprintf(os.Stderr, "Roy (Wikidata): Number of anticipated signatures before: %d\n", before)
	fmt.Fprintf(os.Stderr, "Roy (Wikidata): Number of anticipated signatures after: %d\n", countSignatures())
	fmt.Fprintf(os.Stderr, "Roy (Wikidata): Report generation complete...\n")

	createReportMapping()
}

// WIKIDATA TODO: Another use for ShEX. We need to ensure the data quality is
// consistent so that it can be used reliably. Do we validate via Wikidata and
// do we do it in Roy too to be safe? What other strategies are in use in
// Siegfried.
func isHex(signature string) (bool, error) {
	if len(signature)%2 != 0 {
		return false, fmt.Errorf("Length of HEX is uneven")
	}
	for i := 0; i < len(signature); i += 2 {
		byte := signature[i : i+2]
		_, err := strconv.ParseUint(byte, 16, 64)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
*/
