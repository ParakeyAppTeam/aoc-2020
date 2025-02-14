package main

import "fmt"

func goSum3(list []int) int {
	for i, value1 := range list {
		for j, value2 := range list[i:] {
			for _, value3 := range list[j:] {
				if value1+value2+value3 == 2020 {
					return value1 * value2 * value3
				}
			}
		}
	}
	panic("Crash boom bang")
}

func part_two(input string) {
	list := parseInput(input)
	result := goSum3(list)
	fmt.Println(result)
}
