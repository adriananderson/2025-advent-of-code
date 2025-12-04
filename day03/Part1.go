package day03

import (
	"strconv"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
	defer utils.Timer("3-1")()
	filename := "day03/day03.txt"

	joltage := 0

	if fileContent, err := utils.ReadFileAsLines(filename); err == nil {
		for lineIdx := 0; lineIdx < len(fileContent); lineIdx++ {
			val := findMaxTens(fileContent[lineIdx])
			joltage += val
		}
	}

	return joltage
}

func findMaxTens(line string) int {
	lineLength := len(line)

	tensVal := -1
	tensPos := -1
	onesVal := -1

	//tens
	for index := 0; index < lineLength-1; index++ {
		ch := line[index]
		val, _ := strconv.Atoi(string(ch))
		if val > tensVal {
			tensVal = val
			tensPos = index
		}
	}
	//ones
	for index := tensPos + 1; index < lineLength; index++ {
		ch := line[index]
		val, _ := strconv.Atoi(string(ch))
		if val > onesVal {
			onesVal = val
		}
	}

	return (10*tensVal + onesVal)
}
