package day07

import (
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("7-2")()
	filename := "day07/day07.txt"

	lines, _ := utils.ReadFileAsLines(filename)
	tachyons := map[int]int{}
	start := strings.Index(lines[0], "S")
	tachyons[start] = 1

	splits := 0

	for i := 1; i < len(lines); i++ {
		newTachyons := map[int]int{}
		for tachyonPos, tachyonCount := range tachyons {
			if lines[i][tachyonPos] == '^' {
				splits++
				count, ok := newTachyons[tachyonPos-1]
				if ok {
					newTachyons[tachyonPos-1] = count + tachyonCount
				} else {
					newTachyons[tachyonPos-1] = tachyonCount
				}
				count, ok = newTachyons[tachyonPos+1]
				if ok {
					newTachyons[tachyonPos+1] = count + tachyonCount
				} else {
					newTachyons[tachyonPos+1] = tachyonCount
				}
			} else {
				count, ok := newTachyons[tachyonPos]
				if ok {
					newTachyons[tachyonPos] = count + tachyonCount
				} else {
					newTachyons[tachyonPos] = tachyonCount
				}
			}
		}
		tachyons = newTachyons
	}

	total := 0
	for _, count := range tachyons {
		total += count
	}

	return total
}
