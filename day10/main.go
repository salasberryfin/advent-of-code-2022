package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

const (
	filenameTest = "test-input"
	filename     = "input"
)

//var (
//	// X register
//	register = 1
//)

/*
CPU has a single register: X = 1
Supports only two instructions:
	- addx V: two cycles and X = X + V
	- noop: one cycle
*/

func contains(sub int, slice []int) bool {
	for _, v := range slice {
		if v == sub {
			return true
		}
	}

	return false
}

func getResultPart2(register []int) string {
	var crt string
	var sprite []int

	for i, v := range register {
		if i%40 == 0 {
			crt += "\n"
		}
		sprite = []int{v - 1, v, v + 1}
		if contains(i%40, sprite) {
			crt += "#"
		} else {
			crt += "."
		}
	}

	fmt.Println("CRT line:", crt)

	return crt
}

func getResultPart1(register []int) int {
	var result int
	for i := 20; i <= len(register); i += 40 {
		product := i * register[i-1]
		result += product
	}

	fmt.Printf("Result: %d\n", result)

	return result
}

func calcRegister(input string) []int {
	lines := strings.Split(input, "\n")

	register := 1
	cycles := 0

	var reg []int

	for _, v := range lines {
		splitted := strings.Split(v, " ")
		ins := splitted[0]
		if ins == "addx" {
			val := splitted[1]
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("Converting string to integer: ", err)
			}
			for i := 0; i < 2; i++ {
				reg = append(reg, register)
			}
			register += intVal
		} else if ins == "noop" {
			cycles++
			reg = append(reg, register)
		}
	}

	return reg
}

func main() {
	data := files.ReadFile(filename)

	register := calcRegister(data)
	getResultPart1(register)
	getResultPart2(register)
}
