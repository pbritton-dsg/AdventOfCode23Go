package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var directionMap = map[rune][]int{
	'<': {0, -1},
	'>': {0, 1},
	'^': {-1, 0},
	'v': {1, 0},
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make([][]rune, 0)
	directions := make([]string, 0)
	gettingDirections := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			scanner.Scan()
			gettingDirections = true
		}
		if !gettingDirections {
			row := make([]rune, len(scanner.Text()))
			for i, c := range scanner.Text() {
				row[i] = c
			}
			grid = append(grid, row)
		} else {
			directions = append(directions, scanner.Text())
		}
	}

	fmt.Println(partOne(grid, directions))
}

func partOne(grid [][]rune, directions []string) int {
	defer timeTrack(time.Now(), "partOne")

	row, col := 0, 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '@' {
				row, col = r, c
			}
		}
	}

	for _, directionList := range directions {
		for _, direction := range directionList {
			row, col = moveRobot(grid, row, col, directionMap[direction])
		}
	}

	return countGrid(grid)
}

func moveRobot(grid [][]rune, row, col int, direction []int) (int, int) {
	nr, nc := row+direction[0], col+direction[1]
	canMove := false
	for grid[nr][nc] != '#' {
		if grid[nr][nc] == '.' {
			canMove = true
			break
		}
		nr, nc = nr+direction[0], nc+direction[1]
	}

	if canMove {
		for nr != row || nc != col {
			tempNr, tempNc := nr+(direction[0]*-1), nc+(direction[1]*-1)
			temp := grid[nr][nc]
			grid[nr][nc] = grid[tempNr][tempNc]
			grid[tempNr][tempNc] = temp
			nr, nc = tempNr, tempNc
		}
		return row + direction[0], col + direction[1]
	}

	return row, col
}

func countGrid(grid [][]rune) int {
	sum := 0
	for r, row := range grid {
		for c, col := range row {
			if col == 'O' {
				sum += 100*r + c
			}
		}
	}
	return sum
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
