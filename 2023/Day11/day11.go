package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var rowsWithGalaxies []int
var colsWithGalaxies []int

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]string, 0)
	rNum, galaxy := 0, 1
	for scanner.Scan() {
		row := make([]string, 0)
		line := scanner.Text()
		for col, char := range line {
			val := string(char)
			if val == "#" {
				val = strconv.Itoa(galaxy)
				galaxy += 1
				rowsWithGalaxies = append(rowsWithGalaxies, rNum)
				colsWithGalaxies = append(colsWithGalaxies, col)
			}
			row = append(row, val)
		}
		grid = append(grid, row)
		rNum += 1
	}

	buildMap(grid)
}

func buildMap(grid [][]string) {
	galaxyMap := make(map[string][]int, 0)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != "." {
				galaxyMap[grid[row][col]] = []int{row, col}
			}
		}
	}

	sum := 0
	for key, val := range galaxyMap {
		for key2, val2 := range galaxyMap {
			if key == key2 {
				continue
			}
			//sum += calcDistance(val, val2, grid, 1)
			sum += calcDistance(val, val2, grid, 999999)
		}
		delete(galaxyMap, key)
	}

	fmt.Println(sum)
}

func calcDistance(x, y []int, grid [][]string, multiplier int) int {
	low, high, distance := min(x[0], y[0]), max(x[0], y[0]), 0

	for i := len(grid) - 1; i >= 0; i-- {
		if !slices.Contains(rowsWithGalaxies, i) && i > low && i < high {
			high += multiplier
		}
	}

	distance += high - low
	low, high = min(x[1], y[1]), max(x[1], y[1])

	for i := len(grid[0]) - 1; i >= 0; i-- {
		if !slices.Contains(colsWithGalaxies, i) && i > low && i < high {
			high += multiplier
		}
	}

	return distance + high - low
}
