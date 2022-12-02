package main

import "testing"

var inputLines []string = []string {
	"A Y",
	"B X",
	"C Z",
}

func TestPart1(t *testing.T) {
	result := part1(inputLines)
	if result != 15 {
		t.Errorf("expected part1 to return 15 but got %v.", result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputLines)
	if result != 12 {
		t.Errorf("expected part2 to return 12 but got %v.", result)
	}
}
