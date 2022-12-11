package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

var (
	//filename = "day8-test-input"
	filename = "day8-input"
)

/*
- All trees around the edges are visible
- For tree n
	- if all trees between n and edges are shorter than n -> n is visible
	- only consider trees in the same row or column: up, down, left, right
*/

type neighbors struct {
	Up    []int
	Down  []int
	Left  []int
	Right []int
}

func isVisible(current int, neighbors neighbors) (bool, map[string]int) {
	var visible bool
	view := map[string]int{}

	visible = false

	for k, i := range neighbors.Up {
		if current <= i {
			view["up"]++
			break
		}
		view["up"]++
		if k == len(neighbors.Up)-1 {
			visible = true
		}
	}
	for k, i := range neighbors.Down {
		if current <= i {
			view["down"]++
			break
		}
		view["down"]++
		if k == len(neighbors.Down)-1 {
			visible = true
		}
	}
	for k, i := range neighbors.Left {
		if current <= i {
			//fmt.Printf("It is not visible going left\n")
			view["left"]++
			break
		}
		view["left"]++
		if k == len(neighbors.Left)-1 {
			visible = true
		}
	}
	for k, i := range neighbors.Right {
		if current <= i {
			view["right"]++
			break
		}
		view["right"]++
		if k == len(neighbors.Right)-1 {
			visible = true
		}
	}

	return visible, view
}

func getNeighbors(row, col int, board [][]int) neighbors {
	up, down := row-1, row+1
	left, right := col-1, col+1
	rows, cols := len(board), len(board[0])

	var neighborsMap neighbors

	neighbors := []int{}

	// up
	for i := up; i >= 0; i-- {
		neighbors = append(neighbors, board[i][col])
		neighborsMap.Up = append(neighborsMap.Up, board[i][col])

	}
	// down
	for i := down; i < rows; i++ {
		neighbors = append(neighbors, board[i][col])
		neighborsMap.Down = append(neighborsMap.Down, board[i][col])
	}
	// left
	for i := left; i >= 0; i-- {
		neighbors = append(neighbors, board[row][i])
		neighborsMap.Left = append(neighborsMap.Left, board[row][i])
	}
	// right
	for i := right; i < cols; i++ {
		neighbors = append(neighbors, board[row][i])
		neighborsMap.Right = append(neighborsMap.Right, board[row][i])
	}

	return neighborsMap
}

func main() {
	input := files.ReadFile(filename)

	var mat [][]int
	for _, i := range strings.Split(input, "\n") {
		// rows
		var row []int
		for _, j := range i {
			// columns
			intJ, err := strconv.Atoi(string(j))
			if err != nil {
				log.Fatal("Failed to convert string to integer: ", err)
			}
			row = append(row, intJ)
		}
		mat = append(mat, row)
	}

	// avoid going through col/row edges
	rows, cols := len(mat), len(mat[0])
	visible := 2*cols + 2*rows - 4
	var maxScenicScore int
	for i := 1; i < rows-1; i++ {
		for k := 1; k < cols-1; k++ {
			//fmt.Printf("Node (%d, %d): %d\n", i, k, mat[i][k])
			neighbors := getNeighbors(i, k, mat)
			visibility, views := isVisible(mat[i][k], neighbors)
			if visibility {
				visible++
			}
			treeScenicScore := views["up"] * views["down"] * views["left"] * views["right"]
			if treeScenicScore > maxScenicScore {
				maxScenicScore = treeScenicScore
			}
		}
	}

	fmt.Printf("[PART 1] RESULT:\n\t%d trees are visible\n", visible)
	fmt.Printf("[PART 2] RESULT:\n\tMax scenic score is %d\n", maxScenicScore)
}
