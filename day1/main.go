package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

//var file = "test-input"

var file = "input"

func merge(l, r []int) []int {
	var result []int
	lIndex := 0
	rIndex := 0

	for lIndex < len(l) && rIndex < len(r) {
		if l[lIndex] >= r[rIndex] {
			result = append(result, r[rIndex])
			rIndex++
		} else {
			result = append(result, l[lIndex])
			lIndex++
		}
	}

	for lIndex <= len(l)-1 {
		result = append(result, l[lIndex])
		lIndex++
	}
	for rIndex <= len(r)-1 {
		result = append(result, r[rIndex])
		rIndex++
	}

	return result
}

// can use sort.Ints() from the standard library
// but prefer to implement the algorithm myself
func mergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	middle := len(input) / 2
	left := input[:middle]
	right := input[middle:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func sliceToInt(slice []string) []int {
	var result []int

	for _, v := range slice {
		if len(v) == 0 {
			return result
		}
		intV, err := strconv.Atoi(v)
		if err != nil {
		}
		result = append(result, intV)
	}

	return result
}

func sum(slice []int) int {
	var result int

	for _, v := range slice {
		result += v
	}

	return result
}

func main() {
	content := files.ReadFile(strings.Join([]string{"day1", file}, "-"))
	splitContent := strings.Split(string(content), "\n\n")

	sumCalories := []int{}
	for _, calories := range splitContent {
		intCalories := sliceToInt(strings.Split(calories, "\n"))
		sumCalories = append(sumCalories, sum(intCalories))
	}

	sortedCalories := mergeSort(sumCalories)

	topThree := sortedCalories[len(sortedCalories)-3:]
	log.Println("Three larger calories by elf: ", topThree)
	sumTopThree := sum(topThree)
	log.Println("Sumatory of top three elves: ", sumTopThree)
	log.Println("Most calories: ", sortedCalories[len(sortedCalories)-1])
}
