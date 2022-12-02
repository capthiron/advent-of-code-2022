package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	readInput, _ := os.Open("01/input.txt")
	defer readInput.Close()

	inputScanner := bufio.NewScanner(readInput)
	inputScanner.Split(bufio.ScanLines)

	var inputLines []string
	for inputScanner.Scan() {
		inputLines = append(inputLines, inputScanner.Text())
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
