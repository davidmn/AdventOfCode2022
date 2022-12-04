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
		if processLine(lines[i]) {
			totalOverlaps++
		}
	}

	return totalOverlaps
}

func processLine(line string) bool {
	shifts := strings.Split(line, ",")

	shiftAStart, shiftAEnd := parseShift(shifts[0])
	shiftBStart, shiftBEnd := parseShift(shifts[1])

	shiftARange := utils.GenerateRange(shiftAStart, shiftAEnd)
	shiftBRange := utils.GenerateRange(shiftBStart, shiftBEnd)

	intersection := utils.IntersectionInts(shiftARange, shiftBRange)

	return len(intersection) == len(shiftARange) || len(intersection) == len(shiftBRange)
}

func parseShift(shift string) (int, int) {
	shiftSplit := strings.Split(shift, "-")
	shiftStart, _ := strconv.Atoi(shiftSplit[0])
	shiftEnd, _ := strconv.Atoi(shiftSplit[1])

	return shiftStart, shiftEnd
}