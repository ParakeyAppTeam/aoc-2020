package main

import (
	"fmt"
	"strings"
)

func part_one(input string) {
	fmt.Println(solve_one(parse(input), 3, 1))
}

func parse(input string) (out [][]string) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		out = append(out, strings.Split(line, ""))
	}
	return out
}

func solve_one(slope [][]string, right int, down int) (out int) {
	x, y := 0, 0

	for {
		if y >= len(slope) {
			break
		}

		if slope[y][x] == "#" {
			out++
		}

		x += right
		y += down

		x = x % len(slope[0])
	}

	return out
}
