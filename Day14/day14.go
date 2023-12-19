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
	grid := make([][]rune, 0)
	for scanner.Scan() {
		row := make([]rune, len(scanner.Text()))
		for i, c := range scanner.Text() {
			row[i] = c
		}
		grid = append(grid, row)
	}

	myMap := make(map[int][]int, 0)
	for i := 0; i <= 1000; i++ {
		grid = rollNorth(grid)
		grid = rollWest(grid)
		grid = rollSouth(grid)
		grid = rollEast(grid)

		load := countGrid(grid)

		if val, ok := myMap[load]; !ok {
			myMap[load] = []int{i}
		} else {
			myMap[load] = append(val, i)
		}
	}

	for key, val := range myMap {
		if len(val) > 1 {
			fmt.Println(key, val)
		}
	}
	fmt.Println(0)
}

func countGrid(grid [][]rune) int {
	total := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column] == 79 {
				total += len(grid) - row
			}
		}
	}
	return total
}

func rollNorth(grid [][]rune) [][]rune {
	for c := 0; c < len(grid); c++ {
		colStones := 0
		for r := len(grid[c]) - 1; r >= 0; r-- {
			if grid[r][c] == 79 {
				colStones += 1
				grid[r][c] = 46
			}
			if grid[r][c] == 35 {
				for colStones > 0 {
					grid[r+colStones][c] = 79
					colStones -= 1
				}
			}
			if r == 0 && grid[r][c] == 46 {
				for colStones > 0 {
					grid[r+colStones-1][c] = 79
					colStones -= 1
				}
			}
		}
	}
	return grid
}

func rollSouth(grid [][]rune) [][]rune {
	for c := 0; c < len(grid); c++ {
		colStones := 0
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == 79 {
				colStones += 1
				grid[r][c] = 46
			}
			if grid[r][c] == 35 {
				for colStones > 0 {
					grid[r-colStones][c] = 79
					colStones -= 1
				}
			}
			if r == len(grid)-1 && grid[r][c] == 46 {
				for colStones > 0 {
					grid[r-colStones+1][c] = 79
					colStones -= 1
				}
			}
		}
	}
	return grid
}

func rollEast(grid [][]rune) [][]rune {
	for r := 0; r < len(grid); r++ {
		rowStones := 0
		for c := 0; c < len(grid); c++ {
			if grid[r][c] == 79 {
				rowStones += 1
				grid[r][c] = 46
			}
			if grid[r][c] == 35 {
				for rowStones > 0 {
					grid[r][c-rowStones] = 79
					rowStones -= 1
				}
			}
			if c == len(grid)-1 && grid[r][c] == 46 {
				for rowStones > 0 {
					grid[r][c-rowStones+1] = 79
					rowStones -= 1
				}
			}
		}
	}
	return grid
}

func rollWest(grid [][]rune) [][]rune {
	for r := 0; r < len(grid); r++ {
		rowStones := 0
		for c := len(grid) - 1; c >= 0; c-- {
			if grid[r][c] == 79 {
				rowStones += 1
				grid[r][c] = 46
			}
			if grid[r][c] == 35 {
				for rowStones > 0 {
					grid[r][c+rowStones] = 79
					rowStones -= 1
				}
			}
			if c == 0 && grid[r][c] == 46 {
				for rowStones > 0 {
					grid[r][c+rowStones-1] = 79
					rowStones -= 1
				}
			}
		}
	}
	return grid
}
