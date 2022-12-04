package utils

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	slice := make([]string, 0)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	return slice
}

func Intersection(slice1, slice2 []rune) (inter []rune) {
	hash := make(map[rune]bool)
	for _, e := range slice1 {
		hash[e] = true
	}
	for _, e := range slice2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = removeDuplicates(inter)
	return
}

func removeDuplicates(elements []rune) (deduplicated []rune) {
	encountered := make(map[rune]bool)
	for _, element := range elements {
		if !encountered[element] {
			deduplicated = append(deduplicated, element)
			encountered[element] = true
		}
	}
	return
}


func IntersectionInts(slice1, slice2 []int) (inter []int) {
	hash := make(map[int]bool)
	for _, e := range slice1 {
		hash[e] = true
	}
	for _, e := range slice2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = removeDuplicatesInts(inter)
	return
}

func removeDuplicatesInts(elements []int) (deduplicated []int) {
	encountered := make(map[int]bool)
	for _, element := range elements {
		if !encountered[element] {
			deduplicated = append(deduplicated, element)
			encountered[element] = true
		}
	}
	return
}

func GenerateRange(start, end int) []int {
	r := make([]int, 0)

	for i := start; i <= end; i++ {
		r = append(r, i)
	}

	return r
}