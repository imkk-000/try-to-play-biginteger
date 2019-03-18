package main

import "fmt"

func ConvertStringToBytes(numberString string) []byte {
	result := []byte{}
	for _, digit := range []rune(numberString) {
		result = append(result, (byte(digit) - '0'))
	}
	return result
}

func main() {
	number := "-12345.67890"
	numberBytes := ConvertStringToBytes(number)
	fmt.Println(numberBytes)
}
