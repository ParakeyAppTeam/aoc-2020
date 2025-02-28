package main

import "fmt"

func part_two(input string) {
	count := 0
	rules, passwords := parse(input)

	for i, password := range passwords {
		if valid2(rules[i], password) {
			count++
		}
	}

	fmt.Println(count)
}

func valid2(rule Rule, text string) bool {

	first := string(text[rule.lower-1]) == rule.target
	second := string(text[rule.upper-1]) == rule.target

	return first != second
}
