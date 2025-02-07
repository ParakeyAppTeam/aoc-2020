package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	result := goSum2020([]int{1721, 979, 366, 299, 675, 1456})
	if result != 514579 {
		t.Errorf("Error test failed!, result: %d", result)
	}
}

func TestParse(t *testing.T) {
	inputString := `1721
979
366
299
675
1456`

	result := parseInput(inputString)

	if !reflect.DeepEqual(result, []int{1721, 979, 366, 299, 675, 1456}) {
		t.Errorf("Error test result! result: %v", result)
	}
}
