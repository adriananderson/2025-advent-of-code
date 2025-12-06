package day06

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
	defer utils.Timer("6-1")()
	filename := "day06/day06.txt"

	lines, _ := utils.ReadFileAsLines(filename)

	lastLine := lines[len(lines)-1]
	operators := strings.Split(lastLine, " ")
	operators = removeEmptyStrings(operators)

	totals := make([]int, len(operators))
	for iii := 0; iii < len(operators); iii++ {
		if operators[iii] == "+" {
			totals[iii] = 0
		} else {
			totals[iii] = 1
		}
	}

	for iii := 0; iii < len(lines)-1; iii++ {
		numbers := strings.Split(lines[iii], " ")
		numbers = removeEmptyStrings(numbers)
		for numberIdx := 0; numberIdx < len(numbers); numberIdx++ {
			number, _ := strconv.Atoi(numbers[numberIdx])
			if operators[numberIdx] == "+" {
				totals[numberIdx] += number
			} else {
				totals[numberIdx] *= number
			}
		}
	}

	total := 0
	for _, totalVal := range totals {
		total += totalVal
	}

	return total
}

func removeEmptyStrings(slice []string) []string {
	var result []string
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
