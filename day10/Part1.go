package day10

import (
	"strconv"
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

type Machine struct {
	lights          []int
	desiredLights   []int
	buttons         [][]int
	joltages        []int
	desiredJoltages []int
	buttonPresses   []int
}

func (m Machine) clone() Machine {
	newMachine := Machine{}
	newMachine.lights = make([]int, len(m.lights))
	copy(newMachine.lights, m.lights)
	newMachine.desiredLights = make([]int, len(m.desiredLights))
	copy(newMachine.desiredLights, m.desiredLights)
	newMachine.buttons = make([][]int, len(m.buttons))
	for iii := 0; iii < len(m.buttons); iii++ {
		newMachine.buttons[iii] = make([]int, len(m.buttons[iii]))
		copy(newMachine.buttons[iii], m.buttons[iii])
	}
	newMachine.buttonPresses = make([]int, len(m.buttonPresses))
	copy(newMachine.buttonPresses, m.buttonPresses)
	newMachine.joltages = make([]int, len(m.joltages))
	copy(newMachine.joltages, m.joltages)
	newMachine.desiredJoltages = make([]int, len(m.desiredJoltages))
	copy(newMachine.desiredJoltages, m.desiredJoltages)
	return newMachine
}

func Part1() int {
	defer utils.Timer("10-1")()
	filename := "day10/day10.txt"
	machines := readMachines(filename)
	//fmt.Println(machines)

	total := 0
	for iii := 0; iii < len(machines); iii++ {
		queue := make([]Machine, 0)
		queue = append(queue, machines[iii])
		fewestClicks := breadthFirstSearch(queue)
		total += fewestClicks
	}

	return total
}

func breadthFirstSearch(machineStates []Machine) int {

	for {
		firstMachineState := machineStates[0]
		machineStates = machineStates[1:]

		for iii := 0; iii < len(firstMachineState.buttons); iii++ { //for each button
			newMachineState := firstMachineState.clone()
			newMachineState.buttonPresses = append(newMachineState.buttonPresses, iii)

			for jjj := 0; jjj < len(firstMachineState.buttons[iii]); jjj++ {

				if newMachineState.lights[newMachineState.buttons[iii][jjj]] == 1 {
					newMachineState.lights[newMachineState.buttons[iii][jjj]] = 0
				} else {
					newMachineState.lights[newMachineState.buttons[iii][jjj]] = 1
				}
			}

			match := true
			for kkk := 0; kkk < len(firstMachineState.lights); kkk++ {
				if newMachineState.lights[kkk] != newMachineState.desiredLights[kkk] {
					match = false
					break
				}
			}
			if match {
				return len(newMachineState.buttonPresses)
			} else {
				machineStates = append(machineStates, newMachineState)
			}
		}
	}
	return -1
}

func readMachines(filename string) []Machine {
	lines, _ := utils.ReadFileAsLines(filename)
	ret := make([]Machine, len(lines))

	for lineIdx := 0; lineIdx < len(lines); lineIdx++ {
		machine := Machine{}

		line := lines[lineIdx]
		line = strings.ReplaceAll(line, "[", "")
		line = strings.ReplaceAll(line, "]", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, "{", "")
		line = strings.ReplaceAll(line, "}", "")

		parts := strings.Split(line, " ")
		lights := make([]int, len(parts[0]))
		desiredLights := make([]int, len(parts[0]))
		machine.lights = lights
		for iii := 0; iii < len(parts[0]); iii++ {
			if parts[0][iii] == '#' {
				desiredLights[iii] = 1
			} else {
				desiredLights[iii] = 0
			}
		}
		machine.desiredLights = desiredLights

		for iii := 1; iii < len(parts)-1; iii++ {
			buttonParts := strings.Split(parts[iii], ",")
			button := make([]int, len(buttonParts))
			for jjj := 0; jjj < len(buttonParts); jjj++ {
				button[jjj], _ = strconv.Atoi(buttonParts[jjj])
			}
			machine.buttons = append(machine.buttons, button)
		}
		machine.buttonPresses = make([]int, 0)

		joltageParts := strings.Split(parts[len(parts)-1], ",")
		machine.joltages = make([]int, len(joltageParts))
		machine.desiredJoltages = make([]int, len(joltageParts))
		for iii := 0; iii < len(joltageParts); iii++ {
			machine.desiredJoltages[iii], _ = strconv.Atoi(joltageParts[iii])
		}

		ret[lineIdx] = machine
	}

	return ret
}
