/*
Starting items: worry leve for each item the monkey holds
Operation: how worry level changes as monkey inspects the item
Test: shows how the monkey uses your worry level to decide where to throw an item next
	- If true: if test == true {}
	- If false: if test == false {}

1. Monkey inspects item
2. Update worry level
3. Monkey gets bored: update worry level
4. Check test for current worry level
5. Item with worry level is thrown to monkey x

Part 2 was hell.
	- The key is to use the Chinese Remainder Theorem
		1. All divisors in from input are prime
		2. Multiply all these divisors to get Greatest Common Divisor (GCD)
		3. Then, before any operation, apply the modulo operation to the worry level with the GCD
		4. Everything else stays the same
Using this method, we maintain the same results to the test of each monkey but drastically reduce the values
we have to manage, which overflow otherwise
*/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

type Monkey struct {
	Items     []int
	Operation string
	Test      int
	True      int
	False     int
	Inspected int
}

const (
	filenameTest = "test-input"
	filename     = "input"
)

func exTurnPart2(monkeys *[]Monkey, divisor int) {
	for i, monkey := range *monkeys {
		for _, item := range monkey.Items {
			(*monkeys)[i].Inspected++
			var worryLevel = item
			operation, strValue := strings.Split(monkey.Operation, " ")[0], strings.Split(monkey.Operation, " ")[1]
			var value int
			worryLevel = worryLevel % divisor
			if strValue == "old" {
				value = worryLevel
			} else {
				value, _ = strconv.Atoi(strValue)
			}
			switch operation {
			case "-":
				worryLevel = worryLevel - value
			case "+":
				worryLevel = worryLevel + value
			case "*":
				worryLevel = worryLevel * value
			case "/":
				worryLevel = worryLevel / value
			}
			var to int
			if worryLevel%monkey.Test != 0 {
				to = monkey.False
			} else {
				to = monkey.True
			}
			(*monkeys)[to].Items = append((*monkeys)[to].Items, worryLevel)
			if len((*monkeys)[i].Items) > 1 {
				(*monkeys)[i].Items = (*monkeys)[i].Items[1:]
			} else {
				(*monkeys)[i].Items = []int{}
			}
		}
	}
}

func exTurnPart1(monkeys *[]Monkey) {
	for i, monkey := range *monkeys {
		for _, item := range monkey.Items {
			(*monkeys)[i].Inspected++
			var worryLevel = item
			//fmt.Printf("Monkey inspects an item with worry level %d\n", worryLevel)
			operation, strValue := strings.Split(monkey.Operation, " ")[0], strings.Split(monkey.Operation, " ")[1]
			var value int
			if strValue == "old" {
				value = worryLevel
			} else {
				value, _ = strconv.Atoi(strValue)
			}
			switch operation {
			case "-":
				worryLevel = item - value
				//fmt.Printf("Worry level decreases by %d to %d\n", value, worryLevel)
			case "+":
				worryLevel = item + value
				//fmt.Printf("Worry level increases by %d to %d\n", value, worryLevel)
			case "*":
				worryLevel = item * value
				//fmt.Printf("Worry level is multiplied by %d to %d\n", value, worryLevel)
			case "/":
				worryLevel = item / value
				//fmt.Printf("Worry level is divided by %d to %d\n", value, worryLevel)
			}
			worryLevel = worryLevel / 3
			var to int
			if worryLevel%monkey.Test != 0 {
				//fmt.Printf("Current worry level is not divisible by %d\n", monkey.Test)
				to = monkey.False
				//fmt.Printf("Item with worry level %d is thrown to monkey %d\n", worryLevel, to)
			} else {
				//fmt.Printf("Current worry level is divisible by %d\n", monkey.Test)
				to = monkey.True
				//fmt.Printf("Item with worry level %d is thrown to monkey %d\n", worryLevel, to)
			}
			(*monkeys)[to].Items = append((*monkeys)[to].Items, worryLevel)
			if len((*monkeys)[i].Items) > 1 {
				(*monkeys)[i].Items = (*monkeys)[i].Items[1:]
			} else {
				(*monkeys)[i].Items = []int{}
			}
		}
	}
}

func main() {
	//data := files.ReadFile(filenameTest)
	data := files.ReadFile(filename)
	var monks = []Monkey{}
	var currentMonkey Monkey
	var divisor = 1
	for _, i := range strings.Split(data, "\n") {
		if i != "" {
			if !strings.HasPrefix(i, "Monkey") {
				prefix := strings.TrimSpace(i)
				strData := strings.Split(strings.TrimSpace(i), ": ")[1]
				if strings.HasPrefix(prefix, "Starting items: ") {
					var details = []int{}
					for _, j := range strings.Split(strData, ", ") {
						v, _ := strconv.Atoi(j)
						details = append(details, v)
					}
					currentMonkey.Items = details
				} else if strings.HasPrefix(prefix, "Operation: ") {
					currentMonkey.Operation = strings.Split(strData, "old ")[1]
				} else if strings.HasPrefix(prefix, "Test: ") {
					v, _ := strconv.Atoi(strings.Split(strData, "by ")[1])
					currentMonkey.Test = v
					divisor *= v
				} else if strings.HasPrefix(prefix, "If true: ") {
					v, _ := strconv.Atoi(strings.Split(strData, "to monkey ")[1])
					currentMonkey.True = v
				} else if strings.HasPrefix(prefix, "If false: ") {
					v, _ := strconv.Atoi(strings.Split(strData, "to monkey ")[1])
					currentMonkey.False = v
				}
			}
		} else {
			monks = append(monks, currentMonkey)
		}
	}
	// append last monkey
	monks = append(monks, currentMonkey)

	var rounds = 10000
	for i := 0; i < rounds; i++ {
		//exTurnPart1(&monks)
		exTurnPart2(&monks, divisor)
	}
	var max1, max2 int
	for i, monk := range monks {
		fmt.Printf("Part1: Monkey %d inspected items %d times\n", i, monk.Inspected)
		if monk.Inspected > max1 {
			old := max1
			max1 = monk.Inspected
			max2 = old
		} else if monk.Inspected > max2 {
			max2 = monk.Inspected
		}
	}
	fmt.Printf("Part 1: The level of monkey business is %d * %d = %d\n", max1, max2, max1*max2)
}
