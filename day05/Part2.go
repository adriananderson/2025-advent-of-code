package day05

import (
	"math"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("5-2")()
	filename := "day05/day05.txt"

	freshIngredients := 0

	ranges, _ := readRanges(filename)

	ranges = mergeRanges(ranges)

	for _, rangePair := range ranges {
		freshIngredients += rangePair.end - rangePair.start + 1
	}

	return freshIngredients
}

func mergeRanges(ranges []RangePair) []RangePair {
	mergedRanges := make([]RangePair, 0)

	//mergeHappened := true
	//for mergeHappened {
	//	mergeHappened := false
	for _, rangePair := range ranges {
		mergeHappened := false
		//find all overlaps with existing ranges if range overlaps existing range
		start := math.MaxInt64
		end := -1

		tempRanges := make([]RangePair, 0)
		for _, pair := range mergedRanges {
			if rangeOverlap(pair, rangePair) {
				mergeHappened = true
				start = min(start, pair.start, rangePair.start)
				end = max(end, pair.end, rangePair.end)
			} else {
				tempRanges = append(tempRanges, pair)
			}
		}
		if mergeHappened {
			mergedRanges = append(tempRanges, RangePair{start, end})
		} else {
			mergedRanges = append(mergedRanges, rangePair)
		}
	}
	//}

	return mergedRanges
}

func rangeOverlap(first RangePair, second RangePair) bool {
	return first.start <= second.end && second.start <= first.end
}
