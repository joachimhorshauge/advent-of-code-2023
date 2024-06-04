package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..\n\n"
	expected := "4361"
	result := part1(input)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..\n\n"
	expected := "467835"
	result := part2(input)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
