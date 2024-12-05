package main

import (
	"aoc24/day01"
	"aoc24/day02"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a day number (1-25)")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid day number")
		return
	}

	switch day {
	case 1:
		fmt.Printf("Day 1 Problem 1:\n%v\n\n", day01.GetProblem1())
		fmt.Printf("Day 1 Solution 1: %d\n\n", day01.Solution1())
		fmt.Printf("Day 1 Problem 2:\n%v\n\n", day01.GetProblem2())
		fmt.Printf("Day 1 Solution 2: %d\n\n", day01.Solution2())
	case 2:
		fmt.Printf("Day 2 Problem 1:\n%v\n\n", day02.GetProblem1())
		fmt.Printf("Day 2 Solution 1: %d\n\n", day02.Solution1())
		fmt.Printf("Day 2 Problem 2:\n%v\n\n", day02.GetProblem2())
		fmt.Printf("Day 2 Solution 2: %d\n\n", day02.Solution2())
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}
}
