package main

import (
	"aoc-2022/lib/input"
	"fmt"
	"log"
	"strings"
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

	x := shape{"X", 1}
	y := shape{"Y", 2}
	z := shape{"Z", 3}

	scoreCalcFn := func(r rule, playerShapeId string) int {
		switch playerShapeId {
		case r.winShape.id:
			return 6 + r.winShape.score
		case r.drawShape.id:
			return 3 + r.drawShape.score
		case r.looseShape.id:
			return 0 + r.looseShape.score
		default:
			return 0
		}
	}

	ruleMap := map[string]rule{
		"A": {y, x, z, scoreCalcFn},
		"B": {z, y, x, scoreCalcFn},
		"C": {x, z, y, scoreCalcFn},
	}

	var score int
	for _, line := range inputLines {

		round := strings.Split(line, " ")
		opponentShapeId := round[0]
		playerShapeId := round[1]

		score += ruleMap[opponentShapeId].getScoredPoints(playerShapeId)
	}

	return score
}

func part2(inputLines []string) int {

	rock := shape{"rock", 1}
	paper := shape{"paper", 2}
	scissors := shape{"scissors", 3}

	scoreCalcFn := func(r rule, playersDestiny string) int {
		switch playersDestiny {
		case "Z":
			return 6 + r.winShape.score
		case "Y":
			return 3 + r.drawShape.score
		case "X":
			return 0 + r.looseShape.score
		default:
			return 0
		}
	}

	ruleMap := map[string]rule{
		"A": {paper, rock, scissors, scoreCalcFn},
		"B": {scissors, paper, rock, scoreCalcFn},
		"C": {rock, scissors, paper, scoreCalcFn},
	}

	var score int
	for _, line := range inputLines {
		round := strings.Split(line, " ")
		opponentShapeId := round[0]
		playersDestiny := round[1]

		score += ruleMap[opponentShapeId].getScoredPoints(playersDestiny)
	}

	return score
}

type shape struct {
	id    string
	score int
}

type rule struct {
	winShape    shape
	drawShape   shape
	looseShape  shape
	scoreCalcFn func(rule, string) int
}

func (r rule) getScoredPoints(value string) int {
	return r.scoreCalcFn(r, value)
}
