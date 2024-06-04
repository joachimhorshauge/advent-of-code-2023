package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	expected := "142"
	result := part1(input)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	expected := "281"
	result := part2(input)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
