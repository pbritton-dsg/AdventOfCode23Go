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

	dfs(grid, row, col, 0)
	return maxDepth / 2
}

func dfs(grid [][]string, row int, col int, depth int) {
	if grid[row][col] == "S" && visited[row][col] == 1 {
		maxDepth = max(maxDepth, depth)
	}
	visited[row][col] = 1
	for key, value := range neighbors {
		fields := strings.Split(key, ",")
		nr, _ := strconv.Atoi(strings.TrimSpace(fields[0]))
		nc, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
		nr += row
		nc += col
		if nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[row]) && ((visited[nr][nc] == 0 && slices.Contains(value, grid[nr][nc])) || grid[nr][nc] == "S") {
			dfs(grid, nr, nc, depth+1)
		}
	}
}
