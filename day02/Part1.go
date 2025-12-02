package day02

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
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
				}
			}
		}
	}

	return invalidIdSum

}
