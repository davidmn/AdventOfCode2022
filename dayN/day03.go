package dayN

import (
	"github.com/davidmn/AdventOfCode2022/utils"
)

func DoTheThing(path string) int {
	file := utils.ReadFile(path)
	return len(file)
}