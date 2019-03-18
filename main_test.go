package main_test

import (
	. "my-biginteger"
	"reflect"
	"testing"
)

func TestConvertStringToBytes(t *testing.T) {
	type TestData struct {
		expectedResult []byte
		inputData      string
	}
	testData := []TestData{
		{[]byte{1, 2, 3, 5}, "1235"},
		{[]byte{253, 1, 2, 3, 5}, "-1235"},
		{[]byte{1, 2, 3, 254, 5}, "123.5"},
		{[]byte{253, 1, 2, 3, 254, 5}, "-123.5"},
	}

	for _, testCase := range testData {
		actualBytes := ConvertStringToBytes(testCase.inputData)
		if !reflect.DeepEqual(testCase.expectedResult, actualBytes) {
			t.Errorf("expect %v but it got %v", testCase.expectedResult, actualBytes)
		}
	}
}
