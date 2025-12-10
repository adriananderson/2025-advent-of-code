package day10

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("10-2")()
	filename := "day10/day10.txt"
	machines := readMachines(filename)

	total := 0
	for iii := 0; iii < len(machines); iii++ {
		fewestClicks := solveWithZ3Python(machines[iii])
		fmt.Printf("%d - Fewest clicks: %d\n", iii, fewestClicks)
		total += fewestClicks
	}

	return total
}

type ProblemData struct {
	NumButtons      int     `json:"num_buttons"`
	NumJoltages     int     `json:"num_joltages"`
	Coefficients    [][]int `json:"coefficients"`
	DesiredJoltages []int   `json:"desired_joltages"`
}

func solveWithZ3Python(machine Machine) int {
	numButtons := len(machine.buttons)
	numJoltages := len(machine.desiredJoltages)

	// Build coefficient matrix
	coefficients := make([][]int, numJoltages)
	for i := 0; i < numJoltages; i++ {
		coefficients[i] = make([]int, numButtons)
	}

	for buttonIdx := 0; buttonIdx < numButtons; buttonIdx++ {
		for _, joltageIdx := range machine.buttons[buttonIdx] {
			coefficients[joltageIdx][buttonIdx]++
		}
	}

	// Create problem data
	problem := ProblemData{
		NumButtons:      numButtons,
		NumJoltages:     numJoltages,
		Coefficients:    coefficients,
		DesiredJoltages: machine.desiredJoltages,
	}

	// Write to JSON file
	jsonData, _ := json.MarshalIndent(problem, "", "  ")
	os.WriteFile("day10/problem.json", jsonData, 0644)

	// Create Python solver script
	pythonScript := `#!/usr/bin/env python3
import json
import sys
from z3 import *

with open('day10/problem.json', 'r') as f:
    data = json.load(f)

num_buttons = data['num_buttons']
num_joltages = data['num_joltages']
coefficients = data['coefficients']
desired_joltages = data['desired_joltages']

# Create optimizer
opt = Optimize()

# Create variables for button presses
buttons = [Int(f'button_{i}') for i in range(num_buttons)]

# Each button press count >= 0
for b in buttons:
    opt.add(b >= 0)

# For each joltage, sum of (button * coefficient) = target
for j in range(num_joltages):
    terms = []
    for b in range(num_buttons):
        if coefficients[j][b] > 0:
            terms.append(buttons[b] * coefficients[j][b])
    
    if terms:
        opt.add(sum(terms) == desired_joltages[j])

# Minimize sum of button presses
opt.minimize(sum(buttons))

# Solve
if opt.check() == sat:
    model = opt.model()
    solution = [model.eval(b).as_long() for b in buttons]
    total = sum(solution)
    print(json.dumps({'total': total, 'presses': solution}))
else:
    print(json.dumps({'total': -1, 'presses': []}))
`

	// Write Python script
	os.WriteFile("day10/solve_z3.py", []byte(pythonScript), 0755)

	// Execute Python script
	fmt.Printf("  Solving with Z3 via Python...\n")
	cmd := exec.Command("python3", "day10/solve_z3.py")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("  Python error: %v\n", err)
		fmt.Printf("  Output: %s\n", string(output))
		return -1
	}

	// Parse result
	var result struct {
		Total   int   `json:"total"`
		Presses []int `json:"presses"`
	}

	if err := json.Unmarshal(output, &result); err != nil {
		fmt.Printf("  JSON parse error: %v\n", err)
		fmt.Printf("  Output: %s\n", string(output))
		return -1
	}

	if result.Total == -1 {
		fmt.Printf("  Z3 found no solution\n")
		return -1
	}

	fmt.Printf("  Z3 solution: %d presses (presses: %v)\n", result.Total, result.Presses)

	// Verify solution
	joltages := make([]int, numJoltages)
	for b := 0; b < numButtons; b++ {
		for j := 0; j < numJoltages; j++ {
			joltages[j] += result.Presses[b] * coefficients[j][b]
		}
	}

	valid := true
	for j := 0; j < numJoltages; j++ {
		if joltages[j] != machine.desiredJoltages[j] {
			valid = false
			fmt.Printf("  WARNING: Joltage %d is %d, expected %d\n", j, joltages[j], machine.desiredJoltages[j])
		}
	}

	if !valid {
		fmt.Printf("  Solution verification FAILED!\n")
		return -1
	}

	return result.Total
}
