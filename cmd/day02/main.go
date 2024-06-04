package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	// Read input file
	data, err := os.ReadFile("cmd/day02/input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s", err)
	}

	// Process the input
	input := strings.TrimSpace(string(data))
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func cubesOfColor(color, line string) int {
	pattern := fmt.Sprintf("(\\d+) %s", color)
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)
	if len(matches) < 2 {
		return 0
	}

	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0
	}
	return num
}

func part1(input string) string {
	// Your Part 1 code here
	lines := strings.Split(input, "\n")
	result := 0

	for _, line := range lines {
		re := regexp.MustCompile(`Game (\d+)`)

		matches := re.FindStringSubmatch(line)
		if len(matches) < 2 {
			continue
		}
		game, _ := strconv.Atoi(matches[1])

		parts := strings.Split(line, ": ")
		if len(parts) < 2 {
			continue
		}

		valid := true
		cubesParts := strings.Split(parts[1], "; ")

		for _, part := range cubesParts {
			if cubesOfColor("red", part) > maxRed ||
				cubesOfColor("green", part) > maxGreen ||
				cubesOfColor("blue", part) > maxBlue {
				valid = false
				break
			}
		}

		if valid {
			result += game
		}
	}

	return fmt.Sprintf("%d", result)
}

func part2(input string) string {
	// Your Part 2 code here
	return "Result of Part 2"
}
