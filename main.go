package main

import (
	"fmt"
	bigInteger "number/mybiginteger"
)

func main() {
	number := "-12345.67890"
	numberBytes := bigInteger.ConvertStringToBytes(number)
	fmt.Println(numberBytes)
}
