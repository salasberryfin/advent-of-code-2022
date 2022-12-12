package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

var (
	//filename = "day9-test-input-part2"
	filename = "day9-test-input"
	//filename="day9-input"
	//moves = map[string]
)

type coords [2]int

type knots [10]coords

func updateVisited(visited []coords, current coords) []coords {
	for _, i := range visited {
		if i == current {
			return visited
		}
	}

	return append(visited, current)
}

func manhattanDistance(old, current coords) int {
	var dist int

	dist = int(math.Abs(float64(old[0]-current[0])) + math.Abs(float64(old[1]-current[1])))

	return dist
}

func moveHead(movement string, head coords) coords {
	headPos := head
	switch movement {
	case "R":
		headPos = coords{headPos[0], headPos[1] + 1}
		//fmt.Printf("head position is %v\n", headPos)
	case "L":
		headPos = coords{headPos[0], headPos[1] - 1}
		//fmt.Printf("head position is %v\n", headPos)
	case "U":
		headPos = coords{headPos[0] + 1, headPos[1]}
		//fmt.Printf("head position is %v\n", headPos)
	case "D":
		headPos = coords{headPos[0] - 1, headPos[1]}
		//fmt.Printf("head position is %v\n", headPos)
	}

	return headPos
}

func updateTail(head, oldHead, tail coords) coords {
	if head[0] != tail[0] && head[1] != tail[1] {
		// only move diagonally if more than two steps away
		if manhattanDistance(tail, head) > 2 {
			tail = oldHead
		}
	} else {
		tail = oldHead
	}

	return tail
}

func main() {
	input := files.ReadFile(filename)
	head := coords{
		0, 0,
	}
	tail := head
	//positions := knots{}
	visited := []coords{tail}
	for _, i := range strings.Split(input, "\n") {
		sliceStr := strings.Split(i, " ")
		direction, numberStr := sliceStr[0], sliceStr[1]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal("Failed to convert string to integer: ", err)
		}
		for k := 0; k < number; k++ {
			oldHead := head
			head = moveHead(direction, head)
			if manhattanDistance(tail, head) > 1 {
				tail = updateTail(head, oldHead, tail)
				visited = updateVisited(visited, tail)
			}
		}
	}
	fmt.Printf("Last head position is %v\n", head)
	fmt.Printf("Tail visited the following nodes %v\n", visited)
	fmt.Printf("Number of visited nodes %d\n", len(visited))
}
