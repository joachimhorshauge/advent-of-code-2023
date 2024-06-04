package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s", err)
	}

	// Process the input
	input := strings.TrimSpace(string(data))
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input string) string {
	// Your Part 1 code here
	return "Result of Part 1"
}

func part2(input string) string {
	// Your Part 2 code here
	return "Result of Part 2"
}
