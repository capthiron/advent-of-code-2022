package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const mainGo = `
package main

import (
	"fmt"
	"log"

	"github.com/capthiron/advent-of-code-2022/pkg/input"
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
	return 1
}

func part2(inputLines []string) int {
	return 1
}
`

const mainTestGo = `
package main

import "testing"

var inputLines []string = []string{}

func TestPart1(t *testing.T) {
	result := part1(inputLines)
	if result != 0 {
		t.Errorf("expected part1 to return 0 but was %v.", result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(inputLines)
	if result != 0 {
		t.Errorf("expected part2 to return 0 but was %v.", result)
	}
}
`

func main() {
	dayFlag := flag.String("day", "", "day you want to bootstrap [01-25]")
	flag.Parse()

	if *dayFlag == "" {
		flag.Usage()
		log.Fatal("no day provided")
	}

	dayHasCorrectFormat, _ := regexp.Match(`^\d\d$`, []byte(*dayFlag)) 
	if !dayHasCorrectFormat {
		flag.Usage()
		log.Fatal("day can only be a value between 01 and 25")
	}

	dayInt, _ := strconv.Atoi(*dayFlag)
	if dayInt < 1 || dayInt > 25 {
		flag.Usage()
		log.Fatal("day is out of range")
	}

	_, err := os.Stat("cmd/" + *dayFlag )
    if !os.IsNotExist(err) {
        log.Fatalf("day %s already exists", *dayFlag)
    }

	//create directory in cmd
	err = os.Mkdir(fmt.Sprintf("cmd/%s", *dayFlag), 0755)
    if err != nil {
        log.Fatal(err)
    }

	//create main.go with part1 and part2 returning part1
	mainFile, err := os.Create(fmt.Sprintf("cmd/%s/main.go", *dayFlag))
	if err != nil {
		log.Fatal(err)
	}
	defer mainFile.Close()
	mainFile.WriteString(mainGo)

	//create main_test.go with TestPart1 and TestPart2 expecting 0
	mainTestFile, err := os.Create(fmt.Sprintf("cmd/%s/main_test.go", *dayFlag))
	if err != nil {
		log.Fatal(err)
	}
	defer mainTestFile.Close()
	mainTestFile.WriteString(mainTestGo)

	//create input.txt
	inputFile, err := os.Create(fmt.Sprintf("cmd/%s/input.txt", *dayFlag))
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
}
