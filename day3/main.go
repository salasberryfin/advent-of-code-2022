package main

import (
	"log"
	"sort"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

var (
	filename = "day3-input"
	//filename = "day3-test-input"

	part1 = false
	part2 = true
)

func getRepeatedItem(left, right []byte) byte {
	lIndex, rIndex := 0, 0
	var result byte

	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] == right[rIndex] {
			result = left[lIndex]
			break
		}
		if left[lIndex] < right[rIndex] {
			lIndex++
		} else {
			rIndex++
		}
	}
	return result
}

func convertCharToValue(compartiment string) byte {
	// a = 97, b = 98... x = 120, y = 121, z = 122
	// A = 65, b = 66... X = 88, Y = 89, Z = 90
	// lower case: priority = byte - 96
	// upper case: priority = byte - 38

	byteComp := []byte(compartiment)
	var priorityValue byte
	for _, v := range byteComp {
		// if upper case character
		if v <= 90 {
			priorityValue = v - 38
		} else {
			// if lower case character
			priorityValue = v - 96
		}
	}

	return priorityValue
}

func getRepeatedItemThree(one, two, three []byte) byte {
	var result byte
	i, j, k := 0, 0, 0

	for i < len(one) && j < len(two) && k < len(three) {
		if one[i] == two[j] && one[i] == three[k] {
			result = one[i]
			break
		} else if one[i] < two[j] {
			i++
		} else if two[j] < three[k] {
			j++
		} else {
			k++
		}
	}

	return result
}

func getElvesBadge(rucksacks []string) int {
	elve1 := strings.Split(rucksacks[0], "")
	elve2 := strings.Split(rucksacks[1], "")
	elve3 := strings.Split(rucksacks[2], "")
	sort.Strings(elve1)
	sort.Strings(elve2)
	sort.Strings(elve3)

	byteElve1 := []byte(strings.Join(elve1, ""))
	byteElve2 := []byte(strings.Join(elve2, ""))
	byteElve3 := []byte(strings.Join(elve3, ""))
	repeated := getRepeatedItemThree(byteElve1, byteElve2, byteElve3)

	value := convertCharToValue(string(repeated))

	return int(value)
}

func getRucksackCompartments(rucksack string) int {
	lenruck := len(rucksack)
	comp1 := strings.Split(rucksack[:lenruck/2], "")
	comp2 := strings.Split(rucksack[lenruck/2:], "")
	sort.Strings(comp1)
	sort.Strings(comp2)

	byteComp1 := []byte(strings.Join(comp1, ""))
	byteComp2 := []byte(strings.Join(comp2, ""))
	repeated := getRepeatedItem(byteComp1, byteComp2)

	value := convertCharToValue(string(repeated))

	return int(value)
}

func main() {
	var score int
	input := files.ReadFile(filename)
	splitContent := strings.Fields(string(input))

	if part1 {
		for _, v := range splitContent {
			if len(v) == 0 {
				break
			}
			score += getRucksackCompartments(v)
		}

		log.Println("[PART 1] Final score is: ", score)
	}

	if part2 {
		var index int
		for index < len(splitContent) {
			elvesRucksacks := splitContent[index : index+3]
			index += 3
			score += getElvesBadge(elvesRucksacks)
		}

		log.Println("[PART 2] Final score is: ", score)
	}
}
