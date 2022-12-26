package main

import (
	"testing"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

func TestExampleInputPart1(t *testing.T) {
	filename := "test-input"
	input := files.ReadFile(filename)

	expected := 13140
	result := getResultPart1(calcRegister(input))

	if result != expected {
		t.Fatalf("Expected %d but got %d\n", expected, result)
	}
}

func TestExampleInputPart2(t *testing.T) {
	filename := "test-input"
	input := files.ReadFile(filename)

	expected := `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	result := getResultPart2(calcRegister(input))

	if result != expected {
		t.Fatalf("Expected %s but got %s\n", expected, result)
	}
}
