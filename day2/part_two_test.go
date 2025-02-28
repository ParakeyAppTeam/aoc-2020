package main

import (
	"testing"
)

func TestPasswordValidation2(t *testing.T) {
	rule := Rule{lower: 1, upper: 3, target: "a"}

	if !valid2(rule, "abc") {
		t.Error("Error test failed!")
	}

	if !valid2(rule, "aabaa") {
		t.Error("Error test failed!")
	}

	if !valid2(rule, "aa") {
		t.Error("Error test failed!")
	}
}

