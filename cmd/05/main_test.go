package main

import "testing"

var inputLines []string = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestPart1(t *testing.T) {
	result := part1(inputLines)
	if result != "CMZ" {
		t.Errorf("expected part1 to return CMZ but was %v.", result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputLines)
	if result != "MCD" {
		t.Errorf("expected part2 to return MCD but was %v.", result)
	}
}
