package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
}

func CountCaloriesDay1Part1() int {
	fileLines := readFile("./inputs/day1.txt")

	var largestTotal = 0
	var total = 0

	for i := 0; i < len(fileLines); i++ {
		if fileLines[i] == "" {
			total = 0
			continue
		}

		calories, _ := strconv.Atoi(fileLines[i])
		total += calories

		if total > largestTotal {
			largestTotal = total
		}
	}

	return largestTotal
}

func CountCaloriesDay1Part2() int {
	fileLines := readFile("./inputs/day1.txt")

	caloriesPerElf := make([]int, 0)

	var total = 0
	for i := 0; i < len(fileLines); i++ {
		if fileLines[i] == "" {
			caloriesPerElf = append(caloriesPerElf, total)
			total = 0
			continue
		}

		calories, _ := strconv.Atoi(fileLines[i])

		total += calories
	}

	sort.Slice(caloriesPerElf, func(p, q int) bool {
		return caloriesPerElf[p] > caloriesPerElf[q]
	})

	return caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2]
}
