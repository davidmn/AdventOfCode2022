package day07

import (
	"math"
	"strconv"
	"strings"

	"github.com/davidmn/AdventOfCode2022/utils"
)

func SolvePart1(path string) int {
	file := utils.ReadFile(path)

	joined := strings.Join(file, "\n")

	commands := strings.Split(joined, "\n$ ")

	currentDirectory := ""
	fileSizes := make(map[string]int)
	directories := make([]string, 0)

	for c := range commands {
		if c == 0 {
			continue // skip root directory
		}

		commandAndOutput := strings.Split(commands[c], "\n")

		commandTokens := strings.Split(commandAndOutput[0], " ")

		switch commandTokens[0] {
		case "ls":
			for _, value := range commandAndOutput[1:] {
				outputTokens := strings.Split(value, " ")
				if outputTokens[0] == "dir" {
					directories = append(directories, currentDirectory+"/"+outputTokens[1])
					continue
				}
				path := currentDirectory + "/" + outputTokens[1]
				size, _ := strconv.Atoi(outputTokens[0])
				fileSizes[path] = size
			}
		case "cd":
			if commandTokens[1] == ".." {
				pathSegments := strings.Split(currentDirectory, "/")
				newPathSegments := pathSegments[1 : len(pathSegments)-1]
				currentDirectory = "/" + strings.Join(newPathSegments, "/")
				if currentDirectory == "/" {
					currentDirectory = ""
				}
			} else {
				currentDirectory += "/" + commandTokens[1]
			}
		}
	}

	directorySizes := make(map[string]int)
	for _, directoryPath := range directories {
		for filePath, fileSize := range fileSizes {
			if strings.HasPrefix(filePath, directoryPath) {
				directorySizes[directoryPath] += fileSize
			}
		}
	}

	total := 0
	for _, size := range directorySizes {
		if size <= 100000 {
			total += size
		}
	}

	return total
}

func SolvePart2(path string) int {
	file := utils.ReadFile(path)

	joined := strings.Join(file, "\n")

	commands := strings.Split(joined, "\n$ ")

	currentDirectory := ""
	fileSizes := make(map[string]int)
	directories := make([]string, 0)

	for c := range commands {
		if c == 0 {
			continue // skip root directory
		}

		commandAndOutput := strings.Split(commands[c], "\n")

		commandTokens := strings.Split(commandAndOutput[0], " ")

		switch commandTokens[0] {
		case "ls":
			for _, value := range commandAndOutput[1:] {
				outputTokens := strings.Split(value, " ")
				if outputTokens[0] == "dir" {
					directories = append(directories, currentDirectory+"/"+outputTokens[1])
					continue
				}
				path := currentDirectory + "/" + outputTokens[1]
				size, _ := strconv.Atoi(outputTokens[0])
				fileSizes[path] = size
			}
		case "cd":
			if commandTokens[1] == ".." {
				pathSegments := strings.Split(currentDirectory, "/")
				newPathSegments := pathSegments[1 : len(pathSegments)-1]
				currentDirectory = "/" + strings.Join(newPathSegments, "/")
				if currentDirectory == "/" {
					currentDirectory = ""
				}
			} else {
				currentDirectory += "/" + commandTokens[1]
			}
		}
	}

	directorySizes := make(map[string]int)
	for _, directoryPath := range directories {
		for filePath, fileSize := range fileSizes {
			if strings.HasPrefix(filePath, directoryPath) {
				directorySizes[directoryPath] += fileSize
			}
		}
	}

	total := 0
	for _, size := range directorySizes {
		if size <= 100000 {
			total += size
		}
	}

	totalUsedSpace := 0
	for _, size := range fileSizes {
		totalUsedSpace += size
	}

	totalDiskSpace := 70000000
	requiredDiskSpace := 30000000
	freeDiskSpace := totalDiskSpace - totalUsedSpace
	spaceNeeded := requiredDiskSpace - freeDiskSpace

	smallestDirectory := math.MaxInt64
	for _, size := range directorySizes {
		if size >= spaceNeeded && size < smallestDirectory {
			smallestDirectory = size
		}
	}

	return smallestDirectory
}
