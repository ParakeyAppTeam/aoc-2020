package main

import (
	"reflect"
	"testing"
)

func TestPasswordValidation(t *testing.T) {
	input := `.#.
#..
.##`

	output := parse(input)

	expect := [][]string{
		{".", "#", "."},
		{"#", ".", "."},
		{".", "#", "#"},
	}

	if !reflect.DeepEqual(expect, output) {
		t.Errorf("Error test result! result: %v", output)
	}

}

func TestTreeCount(t *testing.T) {
	input := [][]string{
		{".", "#", "."},
		{"#", ".", "."},
		{".", "#", "#"},
	}

	output := solve_one(input, 1, 1)

	if output != 1 {
		t.Errorf("Error test result! result: %v", output)
	}
}

func TestWrapAround(t *testing.T) {
	input := [][]string{
		{".", "#", "."},
		{"#", ".", "#"},
		{".", "#", "#"},
	}

	output := solve_one(input, 2, 1)

	if output != 2 {
		t.Errorf("Error test result! result: %v", output)
	}
}
