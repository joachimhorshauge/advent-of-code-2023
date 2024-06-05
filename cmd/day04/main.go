package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Read input file
	data, err := os.ReadFile("cmd/day04/input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s", err)
	}

	// Process the input
	input := strings.TrimSpace(string(data))
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func evalScratchcard(scratchcard []string) int {
	re := regexp.MustCompile(`\d{1,2}`)
	a := re.FindAllString(scratchcard[0], -1)
	b := re.FindAllString(scratchcard[1], -1)
	return int(math.Pow(2, float64(len(intersect.Hash(a, b)))-1))
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	result := 0
	for _, line := range lines {
		line = strings.Split(line, ":")[1]      // Remove prefix
		scratchcard := strings.Split(line, "|") // Remove prefix
		result += evalScratchcard(scratchcard)
	}

	return fmt.Sprintf("%d", result)
}

func part2(input string) string {
	// Your Part 2 code here
	return "Result of Part 2"
}
