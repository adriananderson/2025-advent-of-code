package day08

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

type Box struct {
	id int
	xx int
	yy int
	zz int
}

//type Circuit struct {
//	boxes      []int
//	wireLength float64
//}

type Distance struct {
	aBox     int
	bBox     int
	distance float64
}

func Part1() int {
	defer utils.Timer("8-1")()
	filename := "day08/day08.txt"
	numberOfJoins := 1000

	lines, boxes, circuitToBox, boxToCircuit := readBoxes(filename)

	distances := calculateDistances(lines, boxes)

	//for iii := 0; iii < 10; iii++ {
	//	distance := distances[iii]
	//	aBox := boxes[distance.aBox]
	//	bBox := boxes[distance.bBox]
	//	fmt.Printf("%v %v %f\n", aBox, bBox, distance.distance)
	//}

	for iii := 0; iii < numberOfJoins; iii++ {
		distance := distances[iii]
		aCircuit := boxToCircuit[distance.aBox]
		bCircuit := boxToCircuit[distance.bBox]
		if aCircuit == bCircuit {
			//fmt.Printf("Skipping %d %d\n", aCircuit, bCircuit)
			continue
		}

		//join circuits
		for _, boxId := range circuitToBox[bCircuit] {
			circuitToBox[aCircuit] = append(circuitToBox[aCircuit], boxId)
			boxToCircuit[boxId] = aCircuit
		}
		circuitToBox[bCircuit] = nil
	}

	circuitSizes := make([]int, len(circuitToBox))
	for _, circuit := range circuitToBox {
		if circuit != nil {
			circuitSizes = append(circuitSizes, len(circuit))
		}
	}
	slices.SortFunc(circuitSizes, func(a, b int) int { return b - a })

	total := circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
	return total
}

func calculateDistances(lines []string, boxes map[int]Box) []Distance {
	distances := make([]Distance, 0)

	for iii := 0; iii < len(lines); iii++ {
		aBox := boxes[iii]
		for jjj := iii + 1; jjj < len(lines); jjj++ {
			bBox := boxes[jjj]
			distance := calculateDistance(aBox, bBox)
			distances = append(distances, Distance{aBox.id, bBox.id, distance})
		}
	}
	slices.SortFunc(distances, func(a, b Distance) int { return int(a.distance - b.distance) })
	return distances
}

func readBoxes(filename string) ([]string, map[int]Box, map[int][]int, map[int]int) {
	lines, _ := utils.ReadFileAsLines(filename)

	boxes := make(map[int]Box)
	circuitToBox := make(map[int][]int)
	boxToCircuit := make(map[int]int)

	for lineIdx := 0; lineIdx < len(lines); lineIdx++ {
		line := lines[lineIdx]
		parts := strings.Split(line, ",")
		xx, _ := strconv.ParseInt(parts[0], 10, 32)
		yy, _ := strconv.ParseInt(parts[1], 10, 32)
		zz, _ := strconv.ParseInt(parts[2], 10, 32)
		box := Box{lineIdx, int(xx), int(yy), int(zz)}
		boxes[lineIdx] = box

		boxList := make([]int, 0)
		boxList = append(boxList, lineIdx)
		circuitToBox[lineIdx] = boxList
		boxToCircuit[lineIdx] = lineIdx
	}
	return lines, boxes, circuitToBox, boxToCircuit
}

func calculateDistance(aBox Box, bBox Box) float64 {
	dist := math.Sqrt(math.Pow(float64(aBox.xx-bBox.xx), 2) + math.Pow(float64(aBox.yy-bBox.yy), 2) + math.Pow(float64(aBox.zz-bBox.zz), 2))
	//fmt.Printf("%d %d %f\n", aBox.id, bBox.id, dist)
	return dist
}
