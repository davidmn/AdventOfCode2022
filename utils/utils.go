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
