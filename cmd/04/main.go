package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/capthiron/advent-of-code-2022/pkg/input"
	mapset "github.com/deckarep/golang-set/v2"
)

func main() {

	inputLines, err := input.ReadInput()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Part 1:", part1(inputLines))
	fmt.Println("Part 2:", part2(inputLines))
}

func part1(inputLines []string) int {

	var fullyContainedSections int
	for _, line := range inputLines {
		
		s1, s2 := getCleanupSectionsFrom(line)
		if s1.Difference(s2).Cardinality() == 0 || s2.Difference(s1).Cardinality() == 0 {
			fullyContainedSections++
		}

	}

	return fullyContainedSections
}

func part2(inputLines []string) int {

	var overlappingPairs int
	for _, line := range inputLines {

		s1, s2 := getCleanupSectionsFrom(line)
		if s1.Intersect(s2).Cardinality() > 0 {
			overlappingPairs++
		}

	}

	return overlappingPairs
}

func getCleanupSectionsFrom(line string) (mapset.Set[int], mapset.Set[int]) {
	sections := strings.Split(line, ",")
	s1 := extractSectionFrom(sections[0])
	s2 := extractSectionFrom(sections[1])
	return s1, s2
}

func extractSectionFrom(sectionRange string) mapset.Set[int] {
	start, _ := strconv.Atoi(strings.Split(sectionRange, "-")[0])
	end, _ := strconv.Atoi(strings.Split(sectionRange, "-")[1])

	section := mapset.NewSet[int]()
	for i := start; i <= end; i++ {
		section.Add(i)
	}

	return section
}
