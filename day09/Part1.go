package day09

import (
	"image"
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
	defer utils.Timer("9-1")()
	filename := "day09/day09.txt"

	tiles := readTiles(filename)
	//fmt.Println(tiles)

	maxArea := -1
	for iii := 0; iii < len(tiles)-1; iii++ {
		for jjj := iii + 1; jjj < len(tiles); jjj++ {
			area := calculateArea(tiles[iii], tiles[jjj])
			//fmt.Printf("%d %d %d\n", iii, jjj, area)
			if area > maxArea {
				//fmt.Printf("---------New max: %d\n", area)
				maxArea = area
			}
		}
	}

	return maxArea
}

func calculateArea(aa, bb image.Point) int {
	var xDist int
	if aa.X > bb.X {
		xDist = 1 + aa.X - bb.X
	} else {
		xDist = 1 + bb.X - aa.X
	}
	var yDist int
	if aa.Y > bb.Y {
		yDist = 1 + aa.Y - bb.Y
	} else {
		yDist = 1 + bb.Y - aa.Y
	}
	return xDist * yDist
}

func readTiles(filename string) []image.Point {
	lines, _ := utils.ReadFileAsLines(filename)

	tiles := make([]image.Point, 0)

	for lineIdx := 0; lineIdx < len(lines); lineIdx++ {
		line := lines[lineIdx]
		parts := strings.Split(line, ",")
		xx, _ := strconv.Atoi(parts[0])
		yy, _ := strconv.Atoi(parts[1])
		point := image.Point{xx, yy}

		tiles = append(tiles, point)
	}
	return tiles
}
