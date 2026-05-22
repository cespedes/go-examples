package camel

import (
	"unicode"
)

// CamelCaseToUnderscores converts ThisKindOfString into this_kind_of_string
// It also deals with acronyms:
// - IBAN            -> iban
// - GoogleScholarID -> google_scholar_id
// - EUCountries     -> eu_countries
func CamelCaseToUnderscores(s string) string {
	b := make([]rune, 0, 2*len(s))

	var last rune
	var acro bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if last != 0 {
				if !acro {
					if len(b) > 0 {
						b = append(b, '_')
					}
					acro = true
				}
				b = append(b, last)
			}
			last = unicode.ToLower(ch)
		} else {
			if last != 0 {
				if len(b) > 0 {
					b = append(b, '_')
				}
				b = append(b, last)
				last = 0
				acro = false
			}
			b = append(b, ch)
		}
	}
	if last != 0 {
		if !acro {
			b = append(b, '_')
		}
		b = append(b, last)
	}
	return string(b)
}
