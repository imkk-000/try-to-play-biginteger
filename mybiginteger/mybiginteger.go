package mybiginteger

import (
	"unicode"
)

func isAllowedRune(index int, r rune) bool {
	return unicode.IsDigit(r) || (index == 0 && r == '-')
}

func ConvertStringToBytes(numberString string) []byte {
	var result []byte
	for index, digit := range []rune(numberString) {
		if !isAllowedRune(index, digit) {
			return nil
		}
		result = append(result, (byte(digit) - '0'))
	}
	return result
}

func Add(number1 []byte, number2 []byte) []byte {
	number1Length := len(number1)
	number2Length := len(number2)
	if number1Length == 0 {
		return nil
	}
	if number2Length == 0 {
		return nil
	}
	diffLength := number1Length - number2Length
	if diffLength < 0 {
		number1, number2 = number2, number1
		diffLength *= -1
	}
	number2 = append(make([]byte, diffLength), number2...)
	number1Length = len(number1)
	number2Length = len(number2)

	accumulatorBytes := make([]byte, number1Length+1)
	var accumulatorPerDigit byte
	for index := number1Length; index > 0; index-- {
		accumulatorPerDigit = (number1[index-1] + number2[index-1])
		accumulatorBytes[index] += accumulatorPerDigit % 10
		accumulatorBytes[index-1] = accumulatorPerDigit / 10
	}
	if accumulatorBytes[0] == 0 {
		return accumulatorBytes[1:]
	}
	return accumulatorBytes
}
