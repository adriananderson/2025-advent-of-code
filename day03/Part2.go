package day03

import (
	"math"
	"strconv"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("3-2")()
	filename := "day03/day03.txt"

	joltage := 0

	if fileContent, err := utils.ReadFileAsLines(filename); err == nil {
		for lineIdx := 0; lineIdx < len(fileContent); lineIdx++ {
			val := findMax(fileContent[lineIdx], 12)
			joltage += val
		}
	}

	return joltage
}

func findMax(line string, size int) int {
	lineLength := len(line)

	start := 0
	ret := 0

	for iii := size; iii > 0; iii-- {
		max := -1
		maxPos := -1
		for index := start; index < lineLength-(iii-1); index++ {
			ch := line[index]
			val, _ := strconv.Atoi(string(ch))
			if val > max {
				max = val
				maxPos = index
			}
		}

		start = maxPos + 1
		decimalPower := math.Pow(10, float64(iii-1))
		ret += max * int(decimalPower)
	}

	return ret
}
