package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var neighbors = map[string][]string{
	"-1, 0": {"|", "7", "F"},
	"0, -1": {"-", "L", "F"},
	"0, 1":  {"-", "7", "J"},
	"1, 0":  {"|", "L", "J"},
}

var visited [][]uint8

var maxDepth int = 0

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

	fmt.Println(partOne(grid))
}

func resetVisited(grid [][]string) {
	visited = make([][]uint8, len(grid))
	for i := range grid {
		visited[i] = make([]uint8, len(grid[i]))
	}
}

func partOne(grid [][]string) int {
	resetVisited(grid)
	col, row := -1, -1
	for i, line := range grid {
		col = slices.Index(line, "S")
		if col != -1 {
			row = i
			break
		}
	}

	length := 0
	for key, value := range neighbors {
		fields := strings.Split(key, ",")
		nr, _ := strconv.Atoi(strings.TrimSpace(fields[0]))
		nc, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
		nr += row
		nc += col
		if nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[row]) && slices.Contains(value, grid[nr][nc]) {
			dir := moveDir[key]
			length = traversePath(grid, nr, nc, dir)
			break
		}
	}
	return length / 2
}

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

func getNextMove(moveDir direction, pipe string) (int, int) {
	if moveDir == UP {
		if pipe == "|" {
			return -1, 0
		} else if pipe == "F" {
			return 0, 1
		} else if pipe == "7" {
			return 0, -1
		}
	} else if moveDir == DOWN {
		if pipe == "|" {
			return 1, 0
		} else if pipe == "J" {
			return 0, -1
		} else if pipe == "L" {
			return 0, 1
		}
	} else if moveDir == LEFT {
		if pipe == "-" {
			return 0, -1
		} else if pipe == "F" {
			return 1, 0
		} else if pipe == "L" {
			return -1, 0
		}
	} else if moveDir == RIGHT {
		if pipe == "-" {
			return 0, 1
		} else if pipe == "J" {
			return -1, 0
		} else if pipe == "7" {
			return 1, 0
		}
	}
	return -999, -999
}

func traversePath(grid [][]string, row int, col int, dir direction) int {
	pathLength := 1
	for grid[row][col] != "S" {
		nr, nc := getNextMove(dir, grid[row][col])
		row += nr
		col += nc
		dir = moveDir[strconv.Itoa(nr)+", "+strconv.Itoa(nc)]
		pathLength += 1
	}

	return pathLength
}
