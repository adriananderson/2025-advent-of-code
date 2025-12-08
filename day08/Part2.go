package day08

import (
	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("8-2")()
	filename := "day08/day08.txt"

	lines, boxes, circuitToBox, boxToCircuit := readBoxes(filename)

	distances := calculateDistances(lines, boxes)

	for _, distance := range distances {
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
		delete(circuitToBox, bCircuit)

		if len(circuitToBox) == 1 {
			return boxes[distance.aBox].xx * boxes[distance.bBox].xx
		}
	}

	return -1
}
