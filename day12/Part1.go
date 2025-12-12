package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

type Present struct {
	area int
}

type Region struct {
	xWidth             int
	yWidth             int
	area               int
	neededPresentCount int
	neededPresents     []int
}

func Part1() int {
	defer utils.Timer("12-1")()
	filename := "day12/day12.txt"

	presents, regions := readFile(filename)

	regionCount := 0
	for iii := 0; iii < len(regions); iii++ {
		regionArea := regions[iii].area
		minPresentArea := calcPresentArea(presents, regions[iii].neededPresents)
		boundingBoxArea := regions[iii].neededPresentCount * 9

		if minPresentArea > regionArea {
			//not physically possible
			continue
		}
		if boundingBoxArea <= regionArea {
			regionCount++
			continue
		}
		fmt.Printf("dunno what to do about region %d\n", iii)
	}

	return regionCount
}

func calcPresentArea(presents []Present, neededPresents []int) int {
	totalArea := 0
	for iii, present := range presents {
		totalArea += present.area * neededPresents[iii]
	}
	return totalArea
}

func readFile(filename string) (presents []Present, regions []Region) {
	presents = []Present{}
	regions = []Region{}

	lines, _ := utils.ReadFileAsLines(filename)
	for lineIdx := 0; lineIdx < len(lines); lineIdx++ {
		if len(lines[lineIdx]) == 0 {
			continue
		}
		if strings.Contains(lines[lineIdx], "x") {
			region := Region{}
			lineParts := strings.Split(lines[lineIdx], ":")
			dimensions := strings.Split(lineParts[0], "x")
			region.xWidth, _ = strconv.Atoi(dimensions[0])
			region.yWidth, _ = strconv.Atoi(dimensions[1])
			region.area = region.xWidth * region.yWidth

			needsString := strings.Split(strings.TrimSpace(lineParts[1]), " ")
			needs := make([]int, len(needsString))
			neededPresentCount := 0
			for iii, needString := range needsString {
				needs[iii], _ = strconv.Atoi(needString)
				neededPresentCount += needs[iii]
			}
			region.neededPresents = needs
			region.neededPresentCount = neededPresentCount

			regions = append(regions, region)
		} else {
			if strings.Contains(lines[lineIdx], ":") {
				continue
			}
			presentIndex := lineIdx / 5
			if presentIndex >= len(presents) {
				presents = append(presents, Present{})
			}
			presents[presentIndex].area += strings.Count(lines[lineIdx], "#")
		}
	}

	return presents, regions
}
