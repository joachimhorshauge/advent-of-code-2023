package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "your test input here"
	expected := "expected output"
	result := part1(input)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := "your test input here"
	expected := "expected output"
	result := part2(input)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
