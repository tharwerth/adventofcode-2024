# Advent of Code 2024

## Project Structure
```
.
├── cmd/
│   └── main.go        # Main runner for all days
├── day01/
│   ├── input.txt      # Day 1 input data
│   ├── problem.txt    # Day 1 problem description
│   └── solution.go    # Day 1 solutions
├── go.mod
└── README.md
```

## Running Solutions

### Run a specific day
```bash
go run cmd/main.go <day>
```

Example:
```bash
go run cmd/main.go 1  # Runs Day 1 solution
go run cmd/main.go 2  # Runs Day 2 solution
```

## Adding New Days
1. Create a new directory `dayXX`
2. Add `solution.go` with `Solution1()` and `Solution2()`
3. Add your `input.txt`, `problem1.txt` and `problem2.txt`
4. Add the new case in `cmd/main.go`
