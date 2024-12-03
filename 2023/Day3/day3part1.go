package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main1() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	grid := buildGrid1(f)

	sum := 0
	for row := 0; row < len(grid); row++ {
		col := 0
		for col < len(grid[row]) {
			numStart := regexp.MustCompile(`\d`).MatchString(grid[row][col])
			num := ""
			include := false
			if numStart {
				for col < len(grid[row]) && regexp.MustCompile(`\d`).MatchString(grid[row][col]) {
					num += grid[row][col]
					if !include {
						include = includeInSum1(row, col, grid)
					}
					col += 1
				}
				if include {
					n, _ := strconv.Atoi(num)
					sum += n
				}
			}
			col += 1
		}
	}

	fmt.Println(sum)
}

func includeInSum1(row int, col int, grid [][]string) bool {
	neighbors := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < len(neighbors); i++ {
		nr, nc := row+neighbors[i][0], col+neighbors[i][1]
		if nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[row]) && !regexp.MustCompile(`\d|\.`).MatchString(grid[nr][nc]) {
			return true
		}

	}
	return false
}

func buildGrid1(f *os.File) [][]string {
	scanner := bufio.NewScanner(f)

	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, r := range line {
			row[i] = string(r)
		}
		grid = append(grid, row)
	}
	return grid
}
