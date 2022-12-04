package dayN

import (
	"strconv"
	"strings"

	"github.com/davidmn/AdventOfCode2022/utils"
)

func SolvePart1(path string) int {
	lines := utils.ReadFile(path)

	totalOverlaps := 0

	for i := 0; i < len(lines); i++ {
		if processLinePart1(lines[i]) {
			totalOverlaps++
		}
	}

	return totalOverlaps
}

func SolvePart2(path string) int {
	lines := utils.ReadFile(path)

	totalOverlaps := 0

	for i := 0; i < len(lines); i++ {
		if processLinePart2(lines[i]) {
			totalOverlaps++
		}
	}

	return totalOverlaps
}

func processLinePart1(line string) bool {
	shifts := strings.Split(line, ",")

	shiftAStart, shiftAEnd := parseShift(shifts[0])
	shiftBStart, shiftBEnd := parseShift(shifts[1])

	shiftARange := utils.GenerateRange(shiftAStart, shiftAEnd)
	shiftBRange := utils.GenerateRange(shiftBStart, shiftBEnd)

	intersection := utils.Intersection(shiftARange, shiftBRange)

	return len(intersection) == len(shiftARange) || len(intersection) == len(shiftBRange)
}

func processLinePart2(line string) bool {
	shifts := strings.Split(line, ",")

	shiftAStart, shiftAEnd := parseShift(shifts[0])
	shiftBStart, shiftBEnd := parseShift(shifts[1])

	shiftARange := utils.GenerateRange(shiftAStart, shiftAEnd)
	shiftBRange := utils.GenerateRange(shiftBStart, shiftBEnd)

	intersection := utils.Intersection(shiftARange, shiftBRange)

	return len(intersection) > 0
}

func parseShift(shift string) (int, int) {
	shiftSplit := strings.Split(shift, "-")
	shiftStart, _ := strconv.Atoi(shiftSplit[0])
	shiftEnd, _ := strconv.Atoi(shiftSplit[1])

	return shiftStart, shiftEnd
}
