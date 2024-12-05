# Advent of Code 2024

## Project Structure

@/
├── cmd/
│ └── main.go # Main runner for all days
├── day01/
│ ├── input.txt
│ ├── prompt.txt
│ └── solution.go
├── go.mod
└── README.md

## Running Solutions

### Run a specific day

```
go run cmd/main.go <day>
```
Example:
```
go run cmd/main.go 1  # Runs Day 1 solution
go run cmd/main.go 2  # Runs Day 2 solution
```

## Adding New Days
1. Create a new directory `dayXX`
2. Add `solution.go` with `SolvePart1()` and `SolvePart2()`
3. Add your `input.txt`
4. Add the new case in `cmd/main.go`
