package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	lower  int
	upper  int
	target string
}

func part_one(input string) {
	count := 0
	rules, passwords := parse(input)

	for i, password := range passwords {
		if valid(rules[i], password) {
			count++
		}
	}

	fmt.Println(count)
}

func parse(input string) ([]Rule, []string) {
	rules := []Rule{}
	passwords := []string{}
	for _, line := range strings.Split(input, "\n") {
		line := strings.Replace(line, ":", "", -1)
		line = strings.Replace(line, "-", " ", -1)
		parts := strings.Split(line, " ")

		lower, _ := strconv.Atoi(parts[0])
		upper, _ := strconv.Atoi(parts[1])

		rules = append(rules, Rule{lower, upper, parts[2]})
		passwords = append(passwords, parts[3])
	}

	return rules, passwords
}

func valid(rule Rule, text string) bool {
	letters := map[string](int){}

	for _, char := range strings.Split(text, "") {
		letters[char] += 1
	}

	return letters[rule.target] <= rule.upper && letters[rule.target] >= rule.lower
}
