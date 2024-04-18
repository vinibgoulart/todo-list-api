package utils

import "unicode"

func ToSnakeCase(s string) string {
	var result string
	for i, rune := range s {
		if unicode.IsUpper(rune) {
			if i > 0 {
				result += "_"
			}
			result += string(unicode.ToLower(rune))
		} else {
			result += string(rune)
		}
	}
	return result
}
