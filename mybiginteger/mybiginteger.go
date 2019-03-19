package mybiginteger

import (
	"unicode"
)

func isAllowedRune(r rune) bool {
	return unicode.IsDigit(r) || (r == '.') || (r == '-')
}

func ConvertStringToBytes(numberString string) []byte {
	result := []byte{}
	for _, digit := range []rune(numberString) {
		if !isAllowedRune(digit) {
			return nil
		}
		result = append(result, (byte(digit) - '0'))
	}
	return result
}
