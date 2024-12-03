package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	grid := buildGrid(f)
	myMap := make(map[string][]int, 0)
	sum := 0
	for row := 0; row < len(grid); row++ {
		col := 0
		for col < len(grid[row]) {
			numStart := regexp.MustCompile(`\d`).MatchString(grid[row][col])
			num := ""
			starPos := ""
			if numStart {
				for col < len(grid[row]) && regexp.MustCompile(`\d`).MatchString(grid[row][col]) {
					num += grid[row][col]
					if starPos == "" {
						starPos = includeInSum(row, col, grid)
					}
					col += 1
				}
				if starPos != "" {
					n, _ := strconv.Atoi(num)
					val, ok := myMap[starPos]
					if ok {
						myMap[starPos] = append(val, n)
					} else {
						myMap[starPos] = []int{n}
					}
				}
			}
			col += 1
		}
	}

	for _, val := range myMap {
		if len(val) == 2 {
			product := 1
			for _, n := range val {
				product *= n
			}
			sum += product
		}
	}

	fmt.Println(sum)
}

func includeInSum(row int, col int, grid [][]string) string {
	neighbors := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < len(neighbors); i++ {
		nr, nc := row+neighbors[i][0], col+neighbors[i][1]
		if nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[row]) && regexp.MustCompile(`\*`).MatchString(grid[nr][nc]) {
			return strconv.Itoa(nr) + "," + strconv.Itoa(nc)
		}

	}
	return ""
}

func buildGrid(f *os.File) [][]string {
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
