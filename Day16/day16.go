package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]string, 0)
	for scanner.Scan() {
		row := make([]string, 0)
		line := scanner.Text()
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	//fmt.Println(partOne(grid))
	fmt.Println(partTwo(grid))
}

func buildEnergized(grid [][]string) [][]int {
	energized := make([][]int, len(grid))
	for r := 0; r < len(energized); r++ {
		energized[r] = make([]int, len(grid[r]))
	}
	return energized
}

func countEnergized() int {
	total := 0
	for r := 0; r < len(energized); r++ {
		for c := 0; c < len(energized[0]); c++ {
			if energized[r][c] == 1 {
				total += 1
			}
		}
	}

	return total
}

func partOne(grid [][]string) int {
	visited := make(map[string]int)
	energized = buildEnergized(grid)
	traversePath(grid, 0, 0, RIGHT, visited)
	return countEnergized()
}

func partTwo(grid [][]string) int {
	max := 0
	c := 0
	for c < len(grid[0]) {
		for _, d := range []direction{DOWN, UP} {
			r := 0
			if d == UP {
				r = len(grid)
			}
			visited := make(map[string]int)
			energized = buildEnergized(grid)
			traversePath(grid, r, c, d, visited)
			total := countEnergized()
			if total > max {
				max = total
			}
		}
		c += 1
	}
	r := 0
	for r < len(grid) {
		for _, d := range []direction{LEFT, RIGHT} {
			c := 0
			if d == RIGHT {
				c = len(grid[0])
			}
			visited := make(map[string]int)
			energized = buildEnergized(grid)
			traversePath(grid, 0, c, d, visited)
			total := countEnergized()
			if total > max {
				max = total
			}
		}
		r += 1
	}
	return max
}

var energized = make([][]int, 0)

type direction int

var moveDir = map[string]direction{
	"-1, 0": UP,
	"0, -1": LEFT,
	"0, 1":  RIGHT,
	"1, 0":  DOWN,
}

const (
	UP    direction = 4
	DOWN            = 3
	LEFT            = 2
	RIGHT           = 1
)

func getNextMove(moveDir direction, pipe string) [][]int {
	nextMoves := make([][]int, 0)
	if moveDir == UP {
		if pipe == "|" || pipe == "." {
			nextMoves = append(nextMoves, []int{-1, 0})
		} else if pipe == "/" {
			nextMoves = append(nextMoves, []int{0, 1})
		} else if pipe == "\\" {
			nextMoves = append(nextMoves, []int{0, -1})
		} else if pipe == "-" {
			nextMoves = append(nextMoves, []int{0, 1})
			nextMoves = append(nextMoves, []int{0, -1})
		}
	} else if moveDir == DOWN {
		if pipe == "|" || pipe == "." {
			nextMoves = append(nextMoves, []int{1, 0})
		} else if pipe == "/" {
			nextMoves = append(nextMoves, []int{0, -1})
		} else if pipe == "\\" {
			nextMoves = append(nextMoves, []int{0, 1})
		} else if pipe == "-" {
			nextMoves = append(nextMoves, []int{0, 1})
			nextMoves = append(nextMoves, []int{0, -1})
		}
	} else if moveDir == LEFT {
		if pipe == "-" || pipe == "." {
			nextMoves = append(nextMoves, []int{0, -1})
		} else if pipe == "\\" {
			nextMoves = append(nextMoves, []int{-1, 0})
		} else if pipe == "/" {
			nextMoves = append(nextMoves, []int{1, 0})
		} else if pipe == "|" {
			nextMoves = append(nextMoves, []int{-1, 0})
			nextMoves = append(nextMoves, []int{1, 0})
		}
	} else if moveDir == RIGHT {
		if pipe == "-" || pipe == "." {
			nextMoves = append(nextMoves, []int{0, 1})
		} else if pipe == "\\" {
			nextMoves = append(nextMoves, []int{1, 0})
		} else if pipe == "/" {
			nextMoves = append(nextMoves, []int{-1, 0})
		} else if pipe == "|" {
			nextMoves = append(nextMoves, []int{-1, 0})
			nextMoves = append(nextMoves, []int{1, 0})
		}
	}
	return nextMoves
}

func traversePath(grid [][]string, row int, col int, dir direction, visited map[string]int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return
	}
	energized[row][col] = 1
	key := strconv.Itoa(row) + "," + strconv.Itoa(col) + "," + strconv.Itoa(int(dir))
	if val, ok := visited[key]; !ok {
		visited[key] = 1
	} else {
		if val > 5 {
			return
		}
		visited[key] = val + 1
	}
	nextMoves := getNextMove(dir, grid[row][col])
	for _, move := range nextMoves {
		dir = moveDir[strconv.Itoa(move[0])+", "+strconv.Itoa(move[1])]
		traversePath(grid, row+move[0], col+move[1], dir, visited)
	}
}
