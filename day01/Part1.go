package day01

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
	defer utils.Timer("1-1")()
	filename := "day01/day01.txt"
	zeroCounts := 0
	position := 50

	if fileContent, err := utils.ReadFileAsLines(filename); err == nil {
		for lineIdx := 0; lineIdx < len(fileContent); lineIdx++ {
			_, afterLeft, leftFound := strings.Cut(fileContent[lineIdx], "L")
			_, afterRight, rightFound := strings.Cut(fileContent[lineIdx], "R")

			if leftFound {
				val, _ := strconv.Atoi(afterLeft)
				position -= val
				for position < 0 {
					position += 100
				}
			}
			if rightFound {
				val, _ := strconv.Atoi(afterRight)
				position += val
				for position > 99 {
					position -= 100
				}
			}
			if position == 0 {
				zeroCounts++
			}
			//fmt.Printf("Line %s: position %d\n", fileContent[lineIdx], position)
		}
	}

	return zeroCounts
}
