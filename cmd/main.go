package main

import (
	"aoc24/day01"
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
		fmt.Printf("Day 1 Problem: %v\n", day01.GetProblem()+"\n")
		fmt.Printf("Day 1 Solution: %v\n", day01.Solution())
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}
}
