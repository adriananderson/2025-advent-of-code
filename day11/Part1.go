package day11

import (
	"strings"

	"github.com/adriananderson/2025-advent-of-code/utils"
)

func Part1() int {
	defer utils.Timer("11-1")()
	filename := "day11/day11.txt"

	devices := readDevices(filename)

	total := walkDevices(devices, devices["svr"], "dac")

	return total
}

func readDevices(filename string) map[string][]string {
	devices := map[string][]string{}
	fileContent, _ := utils.ReadFileAsLines(filename)
	for lineIdx := 0; lineIdx < len(fileContent); lineIdx++ {
		line := fileContent[lineIdx]
		keyValuesParts := strings.Split(line, ":")
		key := keyValuesParts[0]
		values := strings.TrimSpace(keyValuesParts[1])
		valueParts := strings.Split(values, " ")

		devices[key] = valueParts
	}
	return devices
}

func walkDevices(devices map[string][]string, currentDevice []string, terminal string) int {
	total := 0
	for _, device := range currentDevice {
		if device == terminal {
			total++
		} else {
			total += walkDevices(devices, devices[device], terminal)
		}
	}
	return total
}
