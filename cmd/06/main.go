package main

import (
	"fmt"
	"log"

	"github.com/capthiron/advent-of-code-2022/pkg/input"
	mapset "github.com/deckarep/golang-set/v2"
)

func main() {

	inputLines, err := input.ReadInput()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Part 1:", part1(inputLines[0]))
	fmt.Println("Part 2:", part2(inputLines[0]))
}

func part1(dataStream string) int {
	markerSeq := dataStream[:4]
	for i := 5; i < len(dataStream); i++ {
		if hasStartOfPacket(markerSeq, 4) {
			break
		}

		markerSeq = dataStream[:i]
	}

	return len(markerSeq)
}

func part2(dataStream string) int {
	markerSeq := dataStream[:14]
	for i := 15; i < len(dataStream); i++ {
		if hasStartOfPacket(markerSeq, 14) {
			break
		}

		markerSeq = dataStream[:i]
	}

	return len(markerSeq)
}

func hasStartOfPacket(sequence string, cardinality int) bool {
	sOPSequence := mapset.NewSet[string]()
	for i := len(sequence) - 1; i >= len(sequence)-cardinality; i-- {
		sOPSequence.Add(string(sequence[i]))
	}

	return sOPSequence.Cardinality() == cardinality
}
