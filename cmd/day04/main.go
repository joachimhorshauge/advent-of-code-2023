package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Scratchcard struct {
	CardNumber int
	A          []string
	B          []string
	amount     int
}

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

func evalScratchcard(card Scratchcard) int {
	return len(intersect.Hash(card.A, card.B))
}

func parseScratchcard(line string) Scratchcard {
	re := regexp.MustCompile(`\d{1,2}`)
	parts := strings.Split(line, ":")
	cardId := parts[0]
	cardNumber, _ := strconv.Atoi(strings.TrimSpace(cardId[5:])) // Remove "Card " prefix

	scratchcard := strings.Split(parts[1], "|")
	a := re.FindAllString(scratchcard[0], -1)
	b := re.FindAllString(scratchcard[1], -1)

	return Scratchcard{
		CardNumber: cardNumber,
		A:          a,
		B:          b,
		amount:     1,
	}
}

func parseAllScratchcards(input string) map[int]Scratchcard {
	lines := strings.Split(input, "\n")
	scratchcards := make(map[int]Scratchcard)

	for _, line := range lines {
		card := parseScratchcard(line)
		scratchcards[card.CardNumber] = card
	}

	return scratchcards
}

func part1(input string) string {
	lines := strings.Split(input, "\n")

	result := 0
	for _, line := range lines {
		card := parseScratchcard(line)
		result += int(math.Pow(2, float64(evalScratchcard(card))-1))
	}

	return fmt.Sprintf("%d", result)
}

func part2(input string) string {
	scratchcards := parseAllScratchcards(input)
	result := 0
	for i := 1; i <= len(scratchcards); i++ {
		result += scratchcards[i].amount
		n := evalScratchcard(scratchcards[i])
		for j := 1; j <= n; j++ {
			if i+j <= len(scratchcards) {
				tempCard := scratchcards[i+j]
				tempCard.amount += scratchcards[i].amount
				scratchcards[i+j] = tempCard
			}
		}
	}
	return fmt.Sprintf("%d", result)
}
