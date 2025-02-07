package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) (result []int) {
	foo := strings.Split(input, "\n")
	for _, value := range foo {
		v, _ := strconv.Atoi(value)
		result = append(result, v)
	}

	return result
}

func goSum2020(list []int) int {
	for i, value1 := range list {
		for _, value2 := range list[i:] {
			if value1+value2 == 2020 {
				return value1 * value2
			}
		}
	}
	panic("Crash boom bang")
}

func part_one(input string) {
	list := parseInput(input)
	result := goSum2020(list)
	fmt.Println(result)
}
