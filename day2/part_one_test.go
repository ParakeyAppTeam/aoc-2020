package main

import (
	"reflect"
	"testing"
)

func TestPasswordValidation(t *testing.T) {
	rule := Rule{lower: 1, upper: 1, target: "a"}

	if !valid(rule, "a") {
		t.Error("Error test failed!")
	}

	if valid(rule, "aabc") {
		t.Error("Error test failed!")
	}

	if valid(rule, "b") {
		t.Error("Error test failed!")
	}
}

func TestParseInput(t *testing.T) {
	text := `1-3 a: abcde
1-3 b: cdef
2-9 c: ccccccccc`

	rules, strings := parse(text)

	if !reflect.DeepEqual(strings, []string{"abcde", "cdef", "ccccccccc"}) {
		t.Errorf("Error test result! result: %v", strings)
	}

	expect := []Rule{{1, 3, "a"}, {1, 3, "b"}, {2, 9, "c"}}
	if !reflect.DeepEqual(rules, expect) {
		t.Errorf("Error test result! result: %v", expect)
	}
}
