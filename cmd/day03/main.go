package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	data, err := os.ReadFile("cmd/day03/input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s", err)
	}

	// Process the input
	input := strings.TrimSpace(string(data))
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func findGears(matrix [][]rune, pattern string) [][]int {
	var gearLocations [][]int
	re := regexp.MustCompile(pattern)

	for i, line := range matrix {
		for j, c := range line {
			if re.MatchString(string(c)) {
				gearLocations = append(gearLocations, []int{i, j})
			}
		}
	}
	return gearLocations
}

func extractNumber(matrix [][]rune, x, y int) int {
	numStr := ""

	// Look left
	for i := y; i >= 0 && matrix[x][i] >= '0' && matrix[x][i] <= '9'; i-- {
		numStr = string(matrix[x][i]) + numStr
	}

	// Look right
	for i := y + 1; i < len(matrix[x]) && matrix[x][i] >= '0' && matrix[x][i] <= '9'; i++ {
		numStr += string(matrix[x][i])
	}

	num, _ := strconv.Atoi(numStr)
	return num
}

func findSurroundingNumbers(matrix [][]rune, gears [][]int, part2 bool) []int {
	var numbers []int
	for _, gear := range gears {
		seenNumbers := make(map[int]bool)
		x, y := gear[0], gear[1]

		surroundings := [][]int{
			{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1},
			{x, y - 1}, {x, y + 1},
			{x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1},
		}

		for _, loc := range surroundings {
			surX, surY := loc[0], loc[1]

			if surX < 0 || surX >= len(matrix) || surY < 0 || surY >= len(matrix[surX]) {
				continue
			}

			char := matrix[surX][surY]
			if char >= '0' && char <= '9' {
				num := extractNumber(matrix, surX, surY)
				if !seenNumbers[num] {
					seenNumbers[num] = true
					if !part2 {
						numbers = append(numbers, num)
					}
				}
			}
		}

		if part2 && len(seenNumbers) >= 2 {
			keys := make([]int, 0, len(seenNumbers))
			for k := range seenNumbers {
				keys = append(keys, k)
			}
			if len(keys) >= 2 {
				product := keys[0] * keys[1]
				numbers = append(numbers, product)
			}
		}
	}
	return numbers
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	gearLocations := findGears(matrix, `[^\.\d]`)
	numbers := findSurroundingNumbers(matrix, gearLocations, false)

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return fmt.Sprintf("%d", sum)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")

	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	gearLocations := findGears(matrix, `\*`)
	numbers := findSurroundingNumbers(matrix, gearLocations, true)

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return fmt.Sprintf("%d", sum)
}
