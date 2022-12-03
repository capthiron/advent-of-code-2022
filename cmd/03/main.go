package main

import (
	"fmt"
	"log"
	"unicode"

	"github.com/capthiron/advent-of-code-2022/pkg/input"
	"github.com/deckarep/golang-set/v2"
)

func main() {

	inputLines, err := input.ReadInput()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Part 1:", part1(inputLines))
	fmt.Println("Part 2:", part2(inputLines))

}

func part1(inputLines []string) int {

	var sum int
	for _, rucksack := range inputLines {

		rucksackItems := []rune(rucksack)
		comp1 := mapset.NewSet(rucksackItems[:len(rucksackItems)/2]...)
		comp2 := mapset.NewSet(rucksackItems[len(rucksackItems)/2:]...)

		for item := range comp1.Intersect(comp2).Iterator().C {
			sum += int(item) - getPriorityFor(item)
		}

	}

	return sum
}

func part2(inputLines []string) int {

	var sum int
	var group = mapset.NewSet[rune]()
	for i, rucksack := range inputLines {

		items := mapset.NewSet[rune]()
		for _, item := range rucksack {
			items.Add(item)
		}

		if group.Cardinality() == 0 {
			group = group.Union(items)
		}
		group = group.Intersect(items)

		if (i+1)%3 == 0 {
			for item := range group.Iterator().C {
				sum += int(item) - getPriorityFor(item)
				break
			}
			group = mapset.NewSet[rune]()
		}

	}

	return sum
}

func getPriorityFor(item rune) int {
	if unicode.IsLower(item) {
		return 96
	}
	return 38
}
