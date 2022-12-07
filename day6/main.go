package main

import (
	"fmt"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

var (
	//filename = "day6-test-input"
	filename = "day6-input"
	part1    = false
	part2    = true
)

func countOccurrences(char, marker string) (bool, int) {
	count := 0
	pos := 0
	for i := 0; i < len(marker); i++ {
		if char == string(marker[i]) {
			count++
			if count > 1 {
				pos = i
			}
		}
	}

	return count > 1, pos
}

func countOccurrencesGo(marker string) bool {
	for _, char := range marker {
		if strings.Count(marker, string(char)) > 1 {
			return true
		}
	}

	return false
}

func findMarker(chars string) int {
	var result int

	var markerLength int
	if part1 {
		markerLength = 4
	} else if part2 {
		markerLength = 14
	}

	for i := 0; i+markerLength <= len(chars); i++ {
		subString := chars[i : i+markerLength]
		if !countOccurrencesGo(subString) {
			return i + markerLength
		}
	}

	return result
}

func main() {
	// get input and analyze characters in groups of <markerLength>
	// find a length = <markerLength> substring that doesn't contain repeated characters
	// return the character number of the last item of the substring
	input := files.ReadFile(filename)

	splitFunc := func(c rune) bool {
		return c == '\n'
	}
	arrInput := strings.FieldsFunc(string(input), splitFunc)

	for _, v := range arrInput {
		fmt.Println("Checking input:", v)
		marker := findMarker(v)
		fmt.Println("\tFinal result:", marker)
	}
}
