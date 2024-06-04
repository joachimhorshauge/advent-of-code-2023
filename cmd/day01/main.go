package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var wordToDigit = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	// Read input file
	data, err := os.ReadFile("cmd/day01/input.txt")
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
	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {
		firstNumber := ""
		lastNumber := ""

		for i := 0; i < len(line); i++ {
			if !unicode.IsDigit(rune(line[i])) {
				continue
			}
			if firstNumber == "" {
				firstNumber = string(line[i])
			}
			lastNumber = string(line[i])
		}
		if firstNumber != "" && lastNumber != "" {
			concatenatedNumber, err := strconv.Atoi(firstNumber + lastNumber)
			if err != nil {
				log.Fatalf("failed to convert concatenated number: %s", err)
			}
			result += concatenatedNumber
		}
	}
	return fmt.Sprintf("%d", result)
}

func part2(input string) string {
	// Your Part 2 code here
	for word := range wordToDigit {
		input = strings.Replace(input, word, word+wordToDigit[word]+word, -1)
	}

	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {
		firstNumber := ""
		lastNumber := ""

		for i := 0; i < len(line); i++ {
			if !unicode.IsDigit(rune(line[i])) {
				continue
			}
			if firstNumber == "" {
				firstNumber = string(line[i])
			}
			lastNumber = string(line[i])
		}
		if firstNumber != "" && lastNumber != "" {
			concatenatedNumber, err := strconv.Atoi(firstNumber + lastNumber)
			if err != nil {
				log.Fatalf("failed to convert concatenated number: %s", err)
			}
			result += concatenatedNumber
		}
	}
	return fmt.Sprintf("%d", result)
}
