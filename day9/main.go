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
	filenameTest2 = "day9-test-input-part2"
	filenameTest1 = "day9-test-input"
	filename      = "day9-input"
)

type coords [2]int

type knots []coords

func (c *coords) add(extraCoords coords) {
	// add two coordinates
	c[0] += extraCoords[0]
	c[1] += extraCoords[1]
}

func updateVisited(visited []coords, current coords) []coords {
	for _, i := range visited {
		if i == current {
			return visited
		}
	}

	return append(visited, current)
}

func moveHeadAdd(movement string, head coords) coords {
	var headPos coords

	switch movement {
	case "R":
		headPos = coords{0, 1}
	case "L":
		headPos = coords{0, -1}
	case "U":
		headPos = coords{1, 0}
	case "D":
		headPos = coords{-1, 0}
	}

	return headPos
}

func updateRope(knot int, rope knots, visited []coords) []coords {
	xDiff, yDiff := rope[knot-1][0]-rope[knot][0], rope[knot-1][1]-rope[knot][1]
	if int(math.Abs(float64(xDiff))) > 1 || int(math.Abs(float64(yDiff))) > 1 {
		if math.Abs(float64(xDiff)) > math.Abs(float64(yDiff)) {
			rope[knot][0] += xDiff - int(xDiff/int(math.Abs(float64(xDiff))))
			rope[knot][1] = rope[knot-1][1]
		} else if math.Abs(float64(yDiff)) > math.Abs(float64(xDiff)) {
			rope[knot][0] = rope[knot-1][0]
			rope[knot][1] += yDiff - int(yDiff/int(math.Abs(float64(yDiff))))
		} else {
			rope[knot][0] += xDiff - int(xDiff/int(math.Abs(float64(xDiff))))
			rope[knot][1] += yDiff - int(yDiff/int(math.Abs(float64(yDiff))))
		}
	}
	return updateVisited(visited, rope[len(rope)-1])
}

func solution(data string, length int) {
	ropes := make(knots, length)
	visited := []coords{ropes[len(ropes)-1]}
	for _, i := range strings.Split(data, "\n") {
		sliceStr := strings.Split(i, " ")
		direction, numberStr := sliceStr[0], sliceStr[1]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal("Failed to convert string to integer: ", err)
		}
		for k := 0; k < number; k++ {
			headMovement := moveHeadAdd(direction, ropes[0])
			ropes[0].add(headMovement)
			for knot := 1; knot < len(ropes); knot++ {
				visited = updateRope(knot, ropes, visited)
			}
		}
	}
	fmt.Println("Visited: ", len(visited))
}

func main() {
	//inputTest1 := files.ReadFile(filenameTest1)
	//solution(inputTest1, 2)
	//input1 := files.ReadFile(filename)
	//solution(input1, 2)
	//inputTest2 := files.ReadFile(filenameTest2)
	//solution(inputTest2, 10)
	input2 := files.ReadFile(filename)
	solution(input2, 10)
}
