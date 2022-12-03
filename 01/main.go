package main

import (
	"aoc-2022/lib/input"
	"fmt"
	"log"
	"sort"
	"strconv"
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

	mostCalories := 0

	caloriesPerElf := 0
	for _, itemCalories := range inputLines {

		itemCaloriesInt, err := strconv.Atoi(itemCalories)
		if err == nil {
			caloriesPerElf += itemCaloriesInt
		}

		if len(itemCalories) == 0 {

			if caloriesPerElf > mostCalories {
				mostCalories = caloriesPerElf
			}

			caloriesPerElf = 0
		}

	}

	if caloriesPerElf != 0 && caloriesPerElf > mostCalories {
		mostCalories = caloriesPerElf
	}

	return mostCalories
}

func part2(inputLines []string) int {

	var caloriesPerElfList []int

	caloriesPerElf := 0
	for _, itemCalories := range inputLines {

		itemCaloriesInt, err := strconv.Atoi(itemCalories)
		if err == nil {
			caloriesPerElf += itemCaloriesInt
		}

		if len(itemCalories) == 0 {
			caloriesPerElfList = append(caloriesPerElfList, caloriesPerElf)
			caloriesPerElf = 0
		}

	}

	if caloriesPerElf != 0 {
		caloriesPerElfList = append(caloriesPerElfList, caloriesPerElf)
	}

	sort.Ints(caloriesPerElfList)

	total := 0
	for _, calories := range caloriesPerElfList[len(caloriesPerElfList)-3:] {
		total += calories
	}

	return total
}
