package main

import "testing"

var inputLinesPart1 = map[string]int{
	"bvwbjplbgvbhsrlpgdmjqwftvncz": 5,
	"nppdvjthqldpwncqszvftbrmjlhg": 6,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw": 11,
}

var inputLinesPart2 = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb": 19,
	"bvwbjplbgvbhsrlpgdmjqwftvncz": 23,
	"nppdvjthqldpwncqszvftbrmjlhg": 23,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw": 26,
}

func TestPart1(t *testing.T) {
	for datastream, markerPosition := range inputLinesPart1 {
		result := part1(datastream)
		if result != markerPosition {
			t.Errorf("expected part1 to return %v for datastream='%s' but was %v.", markerPosition, datastream, result)
		}
	}
}

func TestPart2(t *testing.T) {
	for datastream, markerPosition := range inputLinesPart2 {
		result := part2(datastream)
		if result != markerPosition {
			t.Errorf("expected part2 to return %v for datastream='%s' but was %v.", markerPosition, datastream, result)
		}
	}
}
