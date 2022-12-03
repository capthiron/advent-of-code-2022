package main

import "testing"

var inputLines []string = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestPart1(t *testing.T) {
	result := part1(inputLines)
	if result != 157 {
		t.Errorf("expected part1 to return 157 but got %v.", result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputLines)
	if result != 70 {
		t.Errorf("expected part2 to return 70 but got %v.", result)
	}
}
