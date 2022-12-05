package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/capthiron/advent-of-code-2022/pkg/input"
	stack "github.com/golang-ds/stack/linkedstack"
)

func main() {

	inputLines, err := input.ReadInput()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Part 1:", part1(inputLines))
	fmt.Println("Part 2:", part2(inputLines))
}

func part1(inputLines []string) string {

	crateStacks, rules := parseInput(inputLines)

	for rules.Size() != 0 {
		rule, _ := rules.Pop()
		for i := 0; i < rule.amount; i++ {
			crateToMove, _ := crateStacks[rule.from].Pop()
			crateStacks[rule.to].Push(crateToMove)
		}
	}

	var topCrates string
	for _, crateStack := range crateStacks {
		topCrate, _ := crateStack.Pop()
		topCrates += topCrate
	}

	return topCrates
}

func part2(inputLines []string) int {
	return 1
}

func parseInput(inputLines []string) ([]stack.LinkedStack[string], stack.LinkedStack[rule]) {

	var crateStacks []stack.LinkedStack[string]
	rules := stack.New[rule]()
	var rulesDone bool
	var indexLineDone bool
	for i := len(inputLines) - 1; i >= 0; i-- {
		if rulesDone && !indexLineDone {
			indexLineDone = true
			continue
		}

		line := inputLines[i]

		if line == "" {
			rulesDone = true
			continue
		}

		if rulesDone {

			var cratesInLine []string
			for i := 0; i <= len(line)-3; i += 4 {

				crate := line[i : i+3]
				if len(strings.TrimSpace(crate)) == 0 {
					cratesInLine = append(cratesInLine, "")
					continue
				}

				cratesInLine = append(cratesInLine, string(crate[1]))
			}

			if len(crateStacks) == 0 {
				for range cratesInLine {
					crateStacks = append(crateStacks, stack.New[string]())
				}
			}

			for i, crate := range cratesInLine {
				if crate != "" {
					crateStacks[i].Push(crate)
				}
			}

			continue
		}

		rules.Push(*ruleFrom(line))
	}

	return crateStacks, rules
}

type rule struct {
	amount int
	from   int
	to     int
}

func ruleFrom(line string) *rule {
	re := regexp.MustCompile(`\d+`)
	extraction := re.FindAllString(line, -1)
	amount, _ := strconv.Atoi(extraction[0])
	from, _ := strconv.Atoi(extraction[1])
	to, _ := strconv.Atoi(extraction[2])
	return &rule{amount, from - 1, to - 1}
}
