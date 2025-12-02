package day02

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("2-1")()
	filename := "day02/day02.txt"
	fileContents, _ := utils.ReadFileAsText(filename)
	ranges := strings.Split(fileContents, ",")
	invalidIdSum := 0

	for rangeIdx := 0; rangeIdx < len(ranges); rangeIdx++ {
		rng := ranges[rangeIdx]
		rngParts := strings.Split(rng, "-")
		low, _ := strconv.Atoi(rngParts[0])
		high, _ := strconv.Atoi(rngParts[1])
		for iii := low; iii <= high; iii++ {
			str := strconv.Itoa(iii)
			strLen := len(str)
			if strLen%2 == 0 {
				first := str[:strLen/2]
				second := str[strLen/2:]
				if first == second {
					invalidIdSum += iii
					continue
				}
			}
			if strLen%3 == 0 {
				first := str[:strLen/3]
				second := str[strLen/3 : (2 * strLen / 3)]
				third := str[(2 * strLen / 3):]
				if first == second && first == third {
					invalidIdSum += iii
					continue
				}
			}
			if strLen%4 == 0 {
				first := str[:strLen/4]
				second := str[strLen/4 : (2 * strLen / 4)]
				third := str[(2 * strLen / 4):(3 * strLen / 4)]
				fourth := str[(3 * strLen / 4):]
				if first == second && first == third && first == fourth {
					invalidIdSum += iii
					continue
				}
			}
			if strLen%5 == 0 {
				first := str[:strLen/5]
				second := str[strLen/5 : (2 * strLen / 5)]
				third := str[(2 * strLen / 5):(3 * strLen / 5)]
				fourth := str[(3 * strLen / 5):(4 * strLen / 5)]
				fifth := str[(4 * strLen / 5):]
				if first == second && first == third && first == fourth && first == fifth {
					invalidIdSum += iii
					continue
				}
			}
			if strLen%6 == 0 {
				first := str[:strLen/6]
				second := str[strLen/6 : (2 * strLen / 6)]
				third := str[(2 * strLen / 6):(3 * strLen / 6)]
				fourth := str[(3 * strLen / 6):(4 * strLen / 6)]
				fifth := str[(4 * strLen / 6):(5 * strLen / 6)]
				sixth := str[(5 * strLen / 6):]
				if first == second && first == third && first == fourth && first == fifth && first == sixth {
					invalidIdSum += iii
					continue
				}
			}
			if strLen%7 == 0 {
				first := str[:strLen/7]
				second := str[strLen/7 : (2 * strLen / 7)]
				third := str[(2 * strLen / 7):(3 * strLen / 7)]
				fourth := str[(3 * strLen / 7):(4 * strLen / 7)]
				fifth := str[(4 * strLen / 7):(5 * strLen / 7)]
				sixth := str[(5 * strLen / 7):(6 * strLen / 7)]
				seventh := str[(6 * strLen / 7):]
				if first == second && first == third && first == fourth && first == fifth && first == sixth && first == seventh {
					invalidIdSum += iii
					continue
				}
			}

		}
	}

	return invalidIdSum

}
