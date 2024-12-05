package day01

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Solve the problem by reading the input file, sorting the slices and calculating the total distance
// between pairs and adding them all together
func Solution1() int {
	left, right := parseInput("day01/input.txt")

	sort.Ints(left)
	sort.Ints(right)

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

func Solution2() int {
	left, right := parseInput("day01/input.txt")

	rightCounts := countOccurrences(right)

	sum := 0
	for _, num := range left {
		sum += num * rightCounts[num] // If num isn't in rightCounts, it returns 0
	}

	return sum
}

// Distill the problem into a string, the original problem is in problem.txt
func GetProblem1() string {
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

// Distill the problem into a string, the original problem is in problem.txt
func GetProblem2() string {
	return fmt.Sprintf(`
Problem:
--------
For each number in left list, multiply it by how many times it appears in right list.
Example:
Left: 3 4 2 1 3 3
Right: 3 3 3 4

3 appears 3 times (3*3 = 9)
4 appears 1 time  (4*1 = 4)
2 appears 0 times (2*0 = 0)
1 appears 0 times (1*0 = 0)
3 appears 3 times (3*3 = 9)
3 appears 3 times (3*3 = 9)

Sum all products: 9 + 4 + 0 + 0 + 9 + 9 = 31`)
}

// Read the lines from input.txt
func readLines(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n"), nil
}

// Count the occurrences of each number in a list
func countOccurrences(list []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range list {
		counts[num]++
	}
	return counts
}

// Parse the input file into two lists of integers
func parseInput(path string) ([]int, []int) {
	lines, err := readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i, line := range lines {
		nums := strings.Fields(line)
		left[i], _ = strconv.Atoi(nums[0])
		right[i], _ = strconv.Atoi(nums[1])
	}

	return left, right
}
