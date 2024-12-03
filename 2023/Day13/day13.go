package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalPartOne, totalPartTwo := 0, 0
	for scanner.Scan() {
		grid := make([][]rune, 0)
		input := scanner.Text()
		for input != "" {
			row := make([]rune, len(input))
			for i, c := range input {
				row[i] = c
			}
			grid = append(grid, row)
			scanner.Scan()
			input = scanner.Text()
		}
		totalPartOne += countCols(grid, 0) + countRows(grid, 0)
		totalPartTwo += countCols(grid, 1) + countRows(grid, 1)
	}
	fmt.Println("Part 1: ", totalPartOne)
	fmt.Println("Part 2: ", totalPartTwo)
}

func countCols(grid [][]rune, threshold int) int {
	for c := 0; c < len(grid[0])-1; c++ {
		if areColsEqual(grid, c, c+1) <= threshold {
			if checkColMirror(grid, c, c+1, threshold) {
				return c + 1
			}
		}
	}
	return 0
}

func checkColMirror(grid [][]rune, l, r, threshold int) bool {
	errors := 0
	for l >= 0 && r < len(grid[0]) {
		errors += areColsEqual(grid, l, r)
		if errors > threshold {
			return false
		}
		l -= 1
		r += 1
	}
	return errors == threshold
}

func countRows(grid [][]rune, threshold int) int {
	for r := 0; r < len(grid)-1; r++ {
		if areRowsEqual(grid[r], grid[r+1]) <= threshold {
			if checkRowMirror(grid, r, r+1, threshold) {
				return (r + 1) * 100
			}
		}
	}
	return 0
}

func checkRowMirror(grid [][]rune, t, b, threshold int) bool {
	errors := 0
	for t >= 0 && b < len(grid) {
		errors += areRowsEqual(grid[t], grid[b])
		if errors > threshold {
			return false
		}
		t -= 1
		b += 1
	}

	return errors == threshold
}

func areRowsEqual(row1, row2 []rune) int {
	different := 0
	for i := 0; i < len(row1); i++ {
		if row1[i] != row2[i] {
			different += 1
		}
	}
	return different
}

func areColsEqual(grid [][]rune, c1, c2 int) int {
	different := 0
	for r := 0; r < len(grid); r++ {
		if grid[r][c1] != grid[r][c2] {
			different += 1
		}
	}
	return different
}
