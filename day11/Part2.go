package day11

import (
	"fmt"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("11-2")()
	filename := "day11/day11.txt"

	devices := readDevices(filename)

	// Count paths from svr to out that visit both dac and fft
	memo := make(map[string]int)
	result := countPathsWithBothNodes(devices, "svr", "out", false, false, memo)

	return result
}

// Memoization key encodes: current node + whether we've visited dac + whether we've visited fft
func memoKey(node string, visitedDac, visitedFft bool) string {
	return fmt.Sprintf("%s_%t_%t", node, visitedDac, visitedFft)
}

func countPathsWithBothNodes(devices map[string][]string, current, target string, visitedDac, visitedFft bool, memo map[string]int) int {
	// Update visited flags
	if current == "dac" {
		visitedDac = true
	}
	if current == "fft" {
		visitedFft = true
	}

	// Base case: reached target
	if current == target {
		// Only count if we visited both dac and fft
		if visitedDac && visitedFft {
			return 1
		}
		return 0
	}

	// Check memo
	key := memoKey(current, visitedDac, visitedFft)
	if val, ok := memo[key]; ok {
		return val
	}

	// Recursively count paths through all neighbors
	total := 0
	for _, next := range devices[current] {
		total += countPathsWithBothNodes(devices, next, target, visitedDac, visitedFft, memo)
	}

	memo[key] = total
	return total
}
