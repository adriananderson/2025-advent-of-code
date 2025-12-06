package day06

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("6-2")()
	filename := "day06/day06.txt"

	lines, _ := utils.ReadFileAsLines(filename)

	total := 0

	breaks := findBreaks(lines)

	for iii := 1; iii < len(breaks); iii++ {
		start := breaks[iii-1] + 1
		end := breaks[iii]
		//fmt.Printf("%d %d\n", start, end)

		operator := lines[len(lines)-1][start]
		//fmt.Printf("operator: %d\n", operator)
		subTotal := 0
		if operator == '*' {
			subTotal = 1
		}

		numberStrings := make([]string, len(lines)-1)

		for jjj := start; jjj < end; jjj++ {
			for kkk := 0; kkk < len(lines)-1; kkk++ {
				if len(lines[kkk]) > jjj {
					//fmt.Printf("%d %d :: %d\n", kkk, jjj, lines[kkk][jjj])
					numberStrings[jjj-start] = numberStrings[jjj-start] + string(lines[kkk][jjj])
					//fmt.Printf("%s\n", numberStrings[jjj-start])
				}
			}
		}

		for jjj := 0; jjj < len(numberStrings); jjj++ {
			val, _ := strconv.Atoi(strings.TrimSpace(numberStrings[jjj]))
			if val == 0 {
				continue
			}
			if operator == '*' {
				subTotal *= val
			} else {
				subTotal += val
			}
			//fmt.Printf("%d %c ", val, operator)
		}
		//fmt.Printf("= %d\n", subTotal)
		total += subTotal
		//fmt.Printf("\n")
	}

	return total
}

func findBreaks(lines []string) []int {

	lastLine := lines[len(lines)-1]
	maxLineLength := len(lastLine)

	breaks := make([]int, 0)
	breaks = append(breaks, -1)
	for iii := 0; iii < maxLineLength; iii++ {
		spaceCount := 0
		for jjj := 0; jjj < len(lines); jjj++ {
			if len(lines[jjj]) > iii {
				//fmt.Printf("%d %d :: %d\n", jjj, iii, rune(lines[jjj][iii]))
				if lines[jjj][iii] == ' ' {
					spaceCount++
				}
			}
		}
		if spaceCount == len(lines) {
			breaks = append(breaks, iii)
		}
	}

	breaks = append(breaks, maxLineLength+1)

	return breaks
}
