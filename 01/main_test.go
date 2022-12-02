package main

import "testing"

var inputLines []string = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestPart1(t *testing.T) {
	result := part1(inputLines)
	if result != 24000 {
		t.Errorf("expected part1 to return 24000 but was %v.", result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputLines)
	if result != 45000 {
		t.Errorf("expected part2 to return 45000 but was %v.", result)
	}
}
