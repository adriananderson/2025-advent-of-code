package day04

import "github.com/adriananderson/2025-advent-of-code/utils"

func Part2() int {
	defer utils.Timer("4-2")()
	filename := "day04/day04.txt"

	removedRollCount := 0
	grid, _ := readFile(filename)

	for {
		newGrid, removedRolls := removeAvailableRolls(grid)
		if removedRolls == 0 {
			break
		}
		removedRollCount += removedRolls
		grid = newGrid
	}

	return removedRollCount
}

func removeAvailableRolls(grid [][]int) ([][]int, int) {
	newGrid := make([][]int, len(grid))
	removedRolls := 0
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		newGrid[rowIdx] = make([]int, len(grid[0]))
		for colIdx := 0; colIdx < len(grid[0]); colIdx++ {
			if grid[rowIdx][colIdx] == 0 {
				newGrid[rowIdx][colIdx] = 0
			} else {
				count := countNeighbors(grid, rowIdx, colIdx)

				//count Neighbors includes self
				if count <= 4 {
					removedRolls++
					newGrid[rowIdx][colIdx] = 0
				} else {
					newGrid[rowIdx][colIdx] = 1
				}
			}
		}
	}

	return newGrid, removedRolls
}
