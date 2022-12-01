package main

import (
	"bufio"
	"os"
)

func readFile(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	slice := make([]string, 0)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	return slice
}
