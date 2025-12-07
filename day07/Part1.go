package day07

import (
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
	defer utils.Timer("7-1")()
	filename := "day07/day07.txt"

	lines, _ := utils.ReadFileAsLines(filename)
	tachyons := map[int]bool{}
	start := strings.Index(lines[0], "S")
	tachyons[start] = true

	splits := 0

	for i := 1; i < len(lines); i++ {
		newTachyons := map[int]bool{}
		for tachyonPos, _ := range tachyons {
			if lines[i][tachyonPos] == '^' {
				splits++
				newTachyons[tachyonPos-1] = true
				newTachyons[tachyonPos+1] = true
			} else {
				newTachyons[tachyonPos] = true
			}
		}
		tachyons = newTachyons
	}

	return splits
}
