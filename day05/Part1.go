package day05

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

type RangePair struct {
	start int
	end   int
}

func Part1() int {
	defer utils.Timer("5-1")()
	filename := "day05/day05.txt"

	freshIngredients := 0
	ranges, fileContent := readRanges(filename)

	for _, line := range fileContent {
		fresh := false
		if len(line) > 0 && !strings.Contains(line, "-") {
			ingredient, _ := strconv.Atoi(line)
			for _, rangePair := range ranges {
				if ingredient >= rangePair.start && ingredient <= rangePair.end {
					freshIngredients++
					fresh = true
					break
				}
			}
			if fresh {
				continue
			}
		}
	}

	return freshIngredients
}

func readRanges(filename string) ([]RangePair, []string) {
	ranges := make([]RangePair, 0)

	fileContent, _ := utils.ReadFileAsLines(filename)
	for _, line := range fileContent {
		if strings.Contains(line, "-") {
			lineParts := strings.Split(line, "-")
			first, _ := strconv.Atoi(lineParts[0])
			second, _ := strconv.Atoi(lineParts[1])
			newRangePair := RangePair{
				start: first,
				end:   second,
			}
			ranges = append(ranges, newRangePair)
		}
	}
	return ranges, fileContent
}
