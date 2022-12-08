package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

var (
	//filename = "day7-test-input"
	filename = "day7-input"

	commandMarker   = "$"
	sizeLimit       = 100000
	filesystemSpace = 70000000
	emptySpace      = 30000000
)

/*
- Store each path in a n-ary tree
- Get the recursive size of each directory, e.g. '/' contains
all other folders
- Create a sorted array of sizes to select the smallest size that satisfies
the required space
*/

type treeNode struct {
	path     string
	size     int
	children []treeNode
}

func processLs(console []string) (int, int) {
	var output int

	// console are the following n lines printed to the terminal
	i := 0
	for i := 0; i < len(console); i++ {
		consArray := strings.Split(console[i], " ")
		isCommand, _ := detectCommand(console[i])
		if isCommand {
			break
		} else if string(consArray[0]) != "dir" {
			// all non dir files
			intSize, err := strconv.Atoi(consArray[0])
			if err != nil {
				log.Fatal("Failed to convert size string to integer:", err)
			}
			output += intSize
		}
	}

	return output, i + 1
}

func processCd(input []string) string {
	var newPath string

	if len(input) != 2 {
		log.Fatal("Failed to process cd command")
	}
	newPath = input[1]

	return newPath
}

func detectCommand(console string) (bool, string) {
	if string(console[0]) == commandMarker {
		consoleInput := strings.SplitN(console, " ", 2)[1]

		return true, consoleInput
	}

	return false, ""
}

func getRecursiveSize(node *treeNode) int {
	size := node.size
	for i := range node.children {
		size += getRecursiveSize(&node.children[i])
	}

	return size
}

func main() {
	input := files.ReadFile(filename)
	splitFunc := func(c rune) bool {
		return c == '\n'
	}
	arrInput := strings.FieldsFunc(string(input), splitFunc)

	currentPath := ""
	i := 0
	sizesTreeArr := map[string]*treeNode{}
	for i < len(arrInput) {
		isShell, consInput := detectCommand(arrInput[i])
		if isShell {
			command := strings.Split(consInput, " ")[0]
			if command == "cd" {
				currentPath = filepath.Join(currentPath, processCd(strings.Split(consInput, " ")))
				i++
			} else if command == "ls" {
				fileSize, offset := processLs(arrInput[i+1:])
				i = i + offset
				newTreeNode := treeNode{
					path:     currentPath,
					size:     fileSize,
					children: nil,
				}
				for k := range sizesTreeArr {
					if strings.Contains(currentPath, k) {
						sizesTreeArr[k].children = append(sizesTreeArr[k].children, newTreeNode)
					}
				}
				sizesTreeArr[newTreeNode.path] = &newTreeNode
			}
		} else {
			i++
		}
	}

	var recursiveSize, largerFolder int
	var sizesArr []int
	var part1 int
	for _, v := range sizesTreeArr {
		recursiveSize = getRecursiveSize(v)
		sizesArr = append(sizesArr, recursiveSize)
		if recursiveSize > largerFolder {
			largerFolder = recursiveSize
		}
		if recursiveSize < sizeLimit {
			part1 += recursiveSize
		}
	}

	fmt.Println("[PART 1] Final result:", part1)

	usedSpace := largerFolder
	freeSpace := filesystemSpace - usedSpace
	required := emptySpace - freeSpace
	// sort size array to find the smallest item that is larger than required
	sort.Ints(sizesArr)
	for _, v := range sizesArr {
		if v > required {
			fmt.Println("[PART 2] Final result:", v)
			break
		}
	}
}
