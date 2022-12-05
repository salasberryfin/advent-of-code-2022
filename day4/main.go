package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

var (
	//filename = "day4-test-input"
	filename = "day4-input"
	//part     = "part1"
	part = "part2"
)

func checkAnyContained(slc1, slc2 []int) bool {
	if reflect.DeepEqual(slc1, slc2) {
		return true
	}
	i, j := 0, 0
	for i < len(slc1) && j < len(slc2) {
		if slc1[i] == slc2[j] {
			return true
		} else if slc1[i] < slc2[j] {
			i++
		} else if slc2[j] < slc1[i] {
			j++
		}
	}

	return false
}

func checkContained(slc1, slc2 []int) bool {
	// return true if slice is contained
	var candidate, reference []int

	if reflect.DeepEqual(slc1, slc2) {
		return true
	} else if len(slc1) < len(slc2) {
		candidate = slc1
		reference = slc2
	} else {
		candidate = slc2
		reference = slc1
	}

	if candidate[0] >= reference[0] && candidate[len(candidate)-1] <= reference[len(reference)-1] {
		return true
	}

	return false
}

func getIDs(elves []string) [][]int {
	var positions [][]int
	var ids []int
	for _, v := range elves {
		ids = []int{}
		idSlice := strings.Split(v, "-")
		intL, err := strconv.Atoi(idSlice[0])
		if err != nil {
			log.Fatal("Failed to convert ID from string to integer: ", err)
		}
		intR, err := strconv.Atoi(idSlice[1])
		if err != nil {
			log.Fatal("Failed to convert ID from string to integer: ", err)
		}
		for intL <= intR {
			ids = append(ids, intL)
			intL++
		}
		positions = append(positions, ids)
	}

	return positions
}

func main() {
	file, input := files.ReadFileByLine(filename)
	var elves []string
	var score int

	for input.Scan() {
		elves = strings.Split(input.Text(), ",")
		// get positions for each elf
		ids := getIDs(elves)

		if part == "part1" {
			// check if any of the positions array is contained in the other
			if checkContained(ids[0], ids[1]) {
				score++
			}
		}

		if part == "part2" {
			// check if any of the positions array is contained in the other
			if checkAnyContained(ids[0], ids[1]) {
				score++
			}
		}
	}
	log.Println("[", part, "] Final score is:", score)

	file.Close()
}
