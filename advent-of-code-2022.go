package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
}

func CountCaloriesDay1Part1() int {
	file, _ := os.Open("./inputs/day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var largestTotal = 0
	var total = 0
	for scanner.Scan() {
		var line = scanner.Text()

		if line == "" {
			total = 0
			continue
		}

		calories, _ := strconv.Atoi(line)

		total += calories

		if total > largestTotal {
			largestTotal = total
		}
	}

	return largestTotal
}

func CountCaloriesDay1Part2() int {
	file, _ := os.Open("./inputs/day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	caloriesPerElf := make([]int, 0)

	var total = 0
	for scanner.Scan() {
		var line = scanner.Text()

		if line == "" {
			caloriesPerElf = append(caloriesPerElf, total)
			total = 0
			continue
		}

		calories, _ := strconv.Atoi(line)

		total += calories
	}

	sort.Slice(caloriesPerElf, func(p, q int) bool {
		return caloriesPerElf[p] > caloriesPerElf[q]
	})

	return caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2]
}
