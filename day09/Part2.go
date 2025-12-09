package day09

import (
	"image"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("9-2")()
	filename := "day09/day09.txt"

	tiles := readTiles(filename)
	//fmt.Println(tiles)
	tiles = append(tiles, tiles[0])

	maxArea := -1
	for iii := 0; iii < len(tiles)-1; iii++ {
		for jjj := iii + 1; jjj < len(tiles); jjj++ {
			minX := min(tiles[iii].X, tiles[jjj].X)
			maxX := max(tiles[iii].X, tiles[jjj].X)
			minY := min(tiles[iii].Y, tiles[jjj].Y)
			maxY := max(tiles[iii].Y, tiles[jjj].Y)
			rect := image.Rectangle{image.Point{minX, minY}, image.Point{maxX, maxY}}

			if boxInsidePolygon(rect, tiles) {
				area := calculateArea(tiles[iii], tiles[jjj])
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func boxInsidePolygon(aBox image.Rectangle, vertices []image.Point) bool {
	if !pointInsidePolygon(aBox.Min, vertices) {
		return false
	}
	if !pointInsidePolygon(aBox.Max, vertices) {
		return false
	}
	if !pointInsidePolygon(image.Point{aBox.Min.X, aBox.Max.Y}, vertices) {
		return false
	}
	if !pointInsidePolygon(image.Point{aBox.Max.X, aBox.Min.Y}, vertices) {
		return false
	}

	//check whether any polygon horizontal line intersects box vertical lines
	boxMinX := aBox.Min.X
	boxMinY := aBox.Min.Y
	boxMaxX := aBox.Max.X
	boxMaxY := aBox.Max.Y
	for iii := 1; iii < len(vertices)-1; iii += 2 {
		if vertices[iii].X < boxMinX && vertices[iii+1].X > boxMinX && vertices[iii].Y > boxMinY && vertices[iii].Y < boxMaxY {
			return false
		}
		if vertices[iii].X < boxMaxX && vertices[iii+1].X > boxMaxX && vertices[iii].Y > boxMinY && vertices[iii].Y < boxMaxY {
			return false
		}
	}
	return true
}

func pointInsidePolygon(p image.Point, vertices []image.Point) bool {
	//first and last vertex are the same
	//this may be an unnecessary optimization
	for iii := 0; iii < len(vertices)-1; iii++ {
		if p == vertices[iii] {
			return true
		}
	}

	//first pair of vertices is a vertical line

	//check whether p is between min and max x values
	minX := 999999
	maxX := -1
	for iii := 0; iii < len(vertices)-1; iii += 2 {
		if (vertices[iii].Y <= p.Y && p.Y < vertices[iii+1].Y) || (vertices[iii+1].Y <= p.Y && p.Y < vertices[iii].Y) {
			if vertices[iii].X < minX {
				minX = vertices[iii].X
			}
			if vertices[iii].X > maxX {
				maxX = vertices[iii].X
			}
		}
	}
	if p.X < minX || p.X > maxX {
		return false
	}

	return true
}
