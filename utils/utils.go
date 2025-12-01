package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func ReadFileAsText(path string) (string, error) {
	var fileName string = path
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Cannot read file %s", fileName)

		return "", err
	}
	fileContent := string(file)

	return fileContent, err
}

func ReadFileAsLines(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Cannot read file %s", path)
		return nil, err
	}
	fileContent := string(file)

	return strings.Split(fileContent, "\n"), err
}

func RemoveElementAt[T any](index int, report []T) []T {
	if index < 0 || index >= len(report) {
		return []T{}
	}

	temp := make([]T, 0)
	for idx, value := range report {
		if idx != index {
			temp = append(temp, value)
		}
	}

	return temp
}
