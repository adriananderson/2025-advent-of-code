package day01

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("1-2")()
	filename := "day01/day01.txt"
	zeroCounts := 0
	position := 50

	if fileContent, err := utils.ReadFileAsLines(filename); err == nil {
		for lineIdx := 0; lineIdx < len(fileContent); lineIdx++ {
			_, afterLeft, leftFound := strings.Cut(fileContent[lineIdx], "L")
			_, afterRight, rightFound := strings.Cut(fileContent[lineIdx], "R")
			pointsToZero := 0
			nextPosition := position

			if leftFound {
				val, _ := strconv.Atoi(afterLeft)
				nextPosition -= val
				for nextPosition < 0 {
					nextPosition += 100
				}
				if position == 0 {
					pointsToZero--
				}
				position -= val
				for position < 1 {
					position += 100
					pointsToZero++
				}
			}
			if rightFound {
				val, _ := strconv.Atoi(afterRight)
				nextPosition += val
				for nextPosition > 99 {
					nextPosition -= 100
					pointsToZero++
				}
			}
			position = nextPosition
			//fmt.Printf("Line %s: position %d, zero %d\n", fileContent[lineIdx], position, pointsToZero)
			zeroCounts += pointsToZero
		}
	}

	return zeroCounts
}
