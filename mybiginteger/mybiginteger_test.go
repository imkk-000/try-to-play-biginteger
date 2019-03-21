package mybiginteger_test

import (
	. "number/mybiginteger"
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
		{[]byte(nil), "a1235"},
		{[]byte(nil), "12$35"},
		{[]byte(nil), "1235@"},
	}
	for _, testCase := range testData {
		actualBytes := ConvertStringToBytes(testCase.inputData)
		if !reflect.DeepEqual(testCase.expectedResult, actualBytes) {
			t.Errorf("expect %v but it got %v", testCase.expectedResult, actualBytes)
		}
	}
}

func TestAdd(t *testing.T) {
	type TestData struct {
		expectedResult   []byte
		inputDataNumber1 []byte
		inputDataNumber2 []byte
	}
	testData := []TestData{
		{[]byte{3, 3}, []byte{2, 2}, []byte{1, 1}},
		{[]byte{1, 1, 0}, []byte{1, 1}, []byte{9, 9}},
		{[]byte{2, 0}, []byte{1, 1}, []byte{9}},
		{[]byte{2, 0}, []byte{9}, []byte{1, 1}},
		{[]byte(nil), []byte{}, []byte{1}},
		{[]byte(nil), []byte{1}, []byte{}},
	}
	for _, testCase := range testData {
		actualAddBytes := Add(testCase.inputDataNumber1, testCase.inputDataNumber2)
		if !reflect.DeepEqual(testCase.expectedResult, actualAddBytes) {
			t.Errorf("expect %v but it got %v", testCase.expectedResult, actualAddBytes)
		}
	}
}
