package main

import (
	"fmt"
	"os"
)

func load() string {
	if len(os.Args) < 2 {
		fmt.Println("No input path provided")
		os.Exit(1)
	}

	file := os.Args[1]
	return read(file)
}

func read(path string) string {
	file, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(file)
}
