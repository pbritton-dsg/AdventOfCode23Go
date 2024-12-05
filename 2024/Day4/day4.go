package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var neighbors = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

var nextChar = map[string]string{
	"X": "M",
	"M": "A",
	"A": "S",
	"S": "0",
}

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
	fmt.Println(partTwo(grid))
}

func partOne(grid [][]string) int {
	defer timeTrack(time.Now(), "partOne")
	sum := 0

	for row, a := range grid {
		for col, _ := range a {
			if grid[row][col] == "X" {
				sum += findWords(grid, row, col)
			}
		}
	}

	return sum
}

func findWords(grid [][]string, row, col int) int {
	sum := 0
	for i := 0; i < len(neighbors); i++ {
		currRow, currCol := row, col
		nr, nc := row+neighbors[i][0], col+neighbors[i][1]
		for nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[row]) && nextChar[grid[currRow][currCol]] == grid[nr][nc] || nextChar[grid[currRow][currCol]] == "0" {
			if nextChar[grid[currRow][currCol]] == "0" {
				sum += 1
				break
			}
			currRow, currCol = nr, nc
			nr, nc = currRow+neighbors[i][0], currCol+neighbors[i][1]
		}
	}
	return sum
}

func partTwo(grid [][]string) int {
	defer timeTrack(time.Now(), "partOne")
	sum := 0

	cache := make(map[string]int)

	for r := 0; r < len(grid)-2; r++ {
		for c := 0; c < len(grid[r])-2; c++ {
			center := strconv.Itoa(r+1) + "," + strconv.Itoa(c+1)
			if _, ok := cache[center]; ((grid[r][c] == "M" && grid[r+2][c+2] == "S") || (grid[r][c] == "S" && grid[r+2][c+2] == "M")) &&
				((grid[r+2][c] == "M" && grid[r][c+2] == "S") || (grid[r+2][c] == "S" && grid[r][c+2] == "M")) &&
				grid[r+1][c+1] == "A" &&
				!ok {
				cache[center] = 1
				sum += 1
			}
		}
	}

	fmt.Println(cache)

	return sum
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
