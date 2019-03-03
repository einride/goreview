package lettercase

import "unicode"

func IsCamel(s string) bool {
	for i, r := range s {
		if i == 0 && !unicode.IsUpper(r) {
			return false
		}
		if r == '_' {
			return false
		}
	}
	return true
}
