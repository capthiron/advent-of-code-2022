package main

import "testing"

var inputLines []string = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestPart1(t *testing.T) {
	result := part1(inputLines)
	if result != 2 {
		t.Errorf("expected part1 to return 2 but got %v.", result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputLines)
	if result != 4 {
		t.Errorf("expected part2 to return 4 but got %v.", result)
	}
}
