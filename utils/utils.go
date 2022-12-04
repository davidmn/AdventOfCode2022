package utils

import (
	"bufio"
	"os"
)

type Comparable interface {
	int | rune
}

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

func Intersection[T Comparable](slice1, slice2 []T) (inter []T) {
	hash := make(map[T]bool)
	for _, e := range slice1 {
		hash[e] = true
	}
	for _, e := range slice2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = RemoveDuplicates(inter)
	return
}

func GenerateRange(start, end int) []int {
	r := make([]int, 0)

	for i := start; i <= end; i++ {
		r = append(r, i)
	}

	return r
}

func RemoveDuplicates[T Comparable](elements []T) (deduplicated []T) {
	encountered := make(map[T]bool)
	for _, element := range elements {
		if !encountered[element] {
			deduplicated = append(deduplicated, element)
			encountered[element] = true
		}
	}
	return
}
