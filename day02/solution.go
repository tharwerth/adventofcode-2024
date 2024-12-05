package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution1() int {
	lines, err := readLines("day02/input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}

	safeCount := 0
	for _, line := range lines {
		numbers, err := parseNumbers(line)
		if err != nil {
			fmt.Println("Error parsing numbers:", err)
			continue
		}

		if isSafeReport(numbers) {
			safeCount++
		}
	}

	return safeCount
}

func Solution2() int {
	lines, err := readLines("day02/input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}

	safeCount := 0
	for _, line := range lines {
		numbers, err := parseNumbers(line)
		if err != nil {
			fmt.Println("Error parsing numbers:", err)
			continue
		}

		if isSafeReport(numbers) {
			safeCount++
			continue
		}

		// Try removing each level one by one to see if the report becomes safe
		for i := range numbers {
			withoutI := make([]int, 0, len(numbers)-1)
			withoutI = append(withoutI, numbers[:i]...)
			withoutI = append(withoutI, numbers[i+1:]...)

			if isSafeReport(withoutI) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func GetProblem1() string {
	return `Analyze reactor safety reports. A report is safe if:
1. The levels are either all increasing or all decreasing
2. Any two adjacent levels differ by at least 1 and at most 3
How many reports are safe?`
}

func GetProblem2() string {
	return `The Problem Dampener allows the reactor safety systems to tolerate a single bad level in a report. 
A report is safe if:
1. The levels are either all increasing or all decreasing.
2. Any two adjacent levels differ by at least 1 and at most 3.
3. Removing one level can make an otherwise unsafe report safe.
How many reports are now safe?`
}

// Read the lines from input.txt
func readLines(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n"), nil
}

// Checks if the difference between two adjacent numbers is valid (between 1 and 3)
func isValidDifference(a, b int) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff >= 1 && diff <= 3
}

// Checks if the sequence is consistently increasing or decreasing
func isConsistentlyDirected(numbers []int) bool {
	if len(numbers) <= 1 {
		return true
	}

	isIncreasing := numbers[1] > numbers[0]

	for i := 1; i < len(numbers)-1; i++ {
		currentIncreasing := numbers[i+1] > numbers[i]
		if currentIncreasing != isIncreasing {
			return false
		}
	}

	return true
}

// Checks if a report is safe by validating both rules
func isSafeReport(numbers []int) bool {
	if !isConsistentlyDirected(numbers) {
		return false
	}

	for i := 0; i < len(numbers)-1; i++ {
		if !isValidDifference(numbers[i], numbers[i+1]) {
			return false
		}
	}

	return true
}

// Converts a string of space-separated numbers into a slice of ints
func parseNumbers(line string) ([]int, error) {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))

	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		numbers[i] = num
	}

	return numbers, nil
}
