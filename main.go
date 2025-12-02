package main

import (
	"fmt"
	"time"

	"github.com/adriananderson/2025-advent-of-code/day01"
	"github.com/adriananderson/2025-advent-of-code/day02"
)

func main() {
	start := time.Now()

	//// Day 01
	fmt.Printf("Final result Day 01 part 1: %d\n", day01.Part1())
	fmt.Printf("Final result Day 01 part 2: %d\n", day01.Part2())
	//// Day 02
	fmt.Printf("Final result Day 02 part 1: %d\n", day02.Part1())
	fmt.Printf("Final result Day 01 part 2: %d\n", day02.Part2())

	fmt.Printf("... took %v\n", time.Since(start))
}
