package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

var (
	//filename = "day5-test-input"
	filename = "day5-input"

	part1 = false
	part2 = true
)

func findTop(mat [][]string) string {
	var topString []byte
	for _, v := range mat {
		character := strings.Trim(v[len(v)-1], "[]")
		topString = append(topString, []byte(character)...)
	}

	return string(topString)
}

func parseMove(order string) (int, int, int) {
	var moves, src, dst int

	splitted := strings.Split(order, " from ")

	moves, err := strconv.Atoi(strings.Split(splitted[0], " ")[1])
	src, err = strconv.Atoi(strings.Split(splitted[1], " to ")[0])
	dst, err = strconv.Atoi(strings.Split(splitted[1], " to ")[1])
	if err != nil {
		log.Fatal("Failed to convert string to integer:", err)
	}

	return moves, src - 1, dst - 1
}

func moveCratesPart2(moves []string, crates [][]string) [][]string {
	for _, move := range moves {
		moves, src, dst := parseMove(move)
		endSrc := len(crates[src])
		crates[dst] = append(crates[dst], crates[src][endSrc-moves:endSrc]...)
		crates[src] = crates[src][:endSrc-moves]
	}

	return crates
}

func moveCrates(moves []string, crates [][]string) [][]string {
	for _, move := range moves {
		moves, src, dst := parseMove(move)
		for i := 0; i < moves; i++ {
			crates[dst] = append(crates[dst], crates[src][len(crates[src])-1])
			crates[src] = crates[src][:len(crates[src])-1]
		}
	}

	return crates
}

func transpose(mat [][]string) [][]string {
	transposed := make([][]string, len(mat)+1)

	var column []string
	for _, row := range mat {
		column = append(column, row[0])
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if len(mat[i][j]) > 0 {
				transposed[j] = append(transposed[j], mat[i][j])
			}
		}
	}

	return transposed
}

func reverse(mat [][]string) [][]string {
	reversed := make([][]string, len(mat))

	for x := 0; x < len(mat); x++ {
		array := mat[x]
		for i, j := 0, len(mat[x])-1; i < j; i, j = i+1, j-1 {
			array[i], array[j] = mat[x][j], mat[x][i]
		}
		stringRep := strings.ReplaceAll(strings.Join(array, " "), "-", "")
		array = strings.Split(strings.TrimSpace(stringRep), " ")
		reversed[x] = append(reversed[x], array...)
	}

	return reversed
}

func main() {
	input := files.ReadFile(filename)

	split := strings.Split(string(input), "\n\n")
	stack, rearrangements := split[0], split[1]

	allRows := strings.Split(stack, "\n")
	rows := allRows[:len(allRows)-1]

	var mat [][]string
	for _, v := range rows {
		mat = append(mat, strings.Split(v, " "))
	}
	transposed := transpose(mat)
	reversed := reverse(transposed)

	var resultMat [][]string
	if part1 {
		resultMat = moveCrates(strings.Split(strings.TrimSpace(rearrangements), "\n"), reversed)
	} else if part2 {
		resultMat = moveCratesPart2(strings.Split(strings.TrimSpace(rearrangements), "\n"), reversed)
	}

	result := findTop(resultMat)
	fmt.Println("result:", result)

}
