package day01

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solution() int {
	lines, err := readLines("day01/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Create two slices for left and right numbers
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	// Parse numbers into separate slices
	for i, line := range lines {
		nums := strings.Fields(line)
		left[i], _ = strconv.Atoi(nums[0])
		right[i], _ = strconv.Atoi(nums[1])
	}

	// Sort both slices
	sort.Ints(left)
	sort.Ints(right)

	// Calculate differences
	diff := 0
	for i := range left {
		if left[i] > right[i] {
			diff += left[i] - right[i]
		} else {
			diff += right[i] - left[i]
		}
	}
	return diff
}

// Distill the problem into a string, the original problem is in problem.txt
func GetProblem() string {
	return fmt.Sprintf(`
Problem:
--------
Given pairs of numbers like:
3 7  (distance: 4)
9 3  (distance: 6)
4 4  (distance: 0)

Calculate the sum of all distances between pairs.
Example sum: 4 + 6 + 0 = 10`)
}

// Read the lines from input.txt
func readLines(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n"), nil
}
