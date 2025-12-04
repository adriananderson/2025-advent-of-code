package day04

import (
	"fmt"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
	defer utils.Timer("4-1")()
	filename := "day04/day04.txt"

	accessibleRolls := 0
	grid, _ := readFile(filename)
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[0]); colIdx++ {
			if grid[rowIdx][colIdx] == 1 {
				count := countNeighbors(grid, rowIdx, colIdx)

				//count Neighbors includes self
				if count <= 4 {
					//fmt.Printf("Row %d Col %d has %d neighbors\n", rowIdx, colIdx, count)
					accessibleRolls++
				}
			}
		}
	}

	return accessibleRolls
}

func countNeighbors(grid [][]int, rowIdx int, colIdx int) int {
	count := 0
	for row := rowIdx - 1; row <= rowIdx+1; row++ {
		for col := colIdx - 1; col <= colIdx+1; col++ {
			if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == 1 {
				count++
			}
		}
	}
	return count
}

func readFile(filename string) ([][]int, error) {
	fileContent, _ := utils.ReadFileAsLines(filename)
	grid := make([][]int, len(fileContent))
	for lineIdx := 0; lineIdx < len(fileContent); lineIdx++ {
		grid[lineIdx] = make([]int, len(fileContent[lineIdx]))
		for colIdx := 0; colIdx < len(fileContent[lineIdx]); colIdx++ {
			switch fileContent[lineIdx][colIdx] {
			case '.':
				grid[lineIdx][colIdx] = 0
			case '@':
				grid[lineIdx][colIdx] = 1
			default:
				fmt.Printf("Invalid char %s\n", fileContent[lineIdx][colIdx])
				grid[lineIdx][colIdx] = -1
			}
		}
	}
	return grid, nil
}
