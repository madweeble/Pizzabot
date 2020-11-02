package utils

import (
	"fmt"
	"testing"
)

//func Test_getCoords(t *testing.T) {
//	type args struct {
//		inputString []string
//	}
//	tests := []struct {
//		name string
//		args args
//		want []string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := getCoords(tt.args.inputString); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getCoords() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_getCord(t *testing.T) {
	tests := []struct {
		strValue string
		arrValue []int
	}{
		{"(0, 0)", []int{0,0}},
		{"(1, 3)", []int{1,3}},
		{"(3, 1)", []int{3,1}},
		{"(4, 4)", []int{4,4}},
		{"(4, 4, 4)", nil},
		{"(n, 4)", nil},
		{"(4, n)", nil},
		{"(99, 99)", []int{99,99}},
	}
	for _, tt := range tests {
		array := getCord(tt.strValue)
		if array == nil && tt.arrValue != nil {
			fmt.Printf("ERROR: passed: %v, expected: %v, but got nil\n", tt.strValue, tt.arrValue)
		} else if len(array) != 2 {
			fmt.Printf("ERROR: array is %v elements\n", len(array))
		} else if array[0] != tt.arrValue[0] && array[1] != tt.arrValue[1] {
			fmt.Printf("ERROR: arrays do not match.  Expected %v, got %v.\n", tt.arrValue, array)
		//} else {
		//	fmt.Printf("Expected: %v, got: %v\n", tt.arrValue, array)
		}
	}
}

func Test_validateArgs(t *testing.T) {
	tests := []struct {
		strValue []string
		expected bool
	}{
		{[]string{"5x5"," (1, 3)"," (4, 4)"}, true},
		{[]string{"3x3","(3,1)"}, true},
		{[]string{"1x1","(3)"}, false},
		{[]string{"5","(1, 3)"}, false},
	}
	for _, tt := range tests {
		returnValue := validateArgs(tt.strValue)
		if returnValue != tt.expected {
			fmt.Printf("ERROR: returned value did not match expected value.  Expected %v, got %v.\n", tt.expected, returnValue)
		//} else {
		//	fmt.Printf("Expected: %v, got: %v\n", tt.expected, returnValue)
		}
	}
}