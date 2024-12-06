package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var startChars = map[string][]int{
	"^": {-1, 0},
	">": {0, 1},
	"<": {0, -1},
	"v": {1, 0},
}

var nextDir = map[string][]int{
	"-1,0": {0, 1},
	"0,1":  {1, 0},
	"0,-1": {-1, 0},
	"1,0":  {0, -1},
}

var gridCopy [][]string

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

	//Get start params
	curRow, curCol := 0, 0
	curDir := []int{0, 0}
	for r, row := range grid {
		for c, char := range row {
			if val, ok := startChars[char]; ok {
				curRow, curCol = r, c
				curDir = val
			}
		}
	}

	count := 0
	visited := make(map[string]bool)
	visited[strconv.Itoa(curRow)+","+strconv.Itoa(curCol)] = true
	nr, nc := curRow+curDir[0], curCol+curDir[1]
	for nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[0]) {
		for grid[nr][nc] == "#" {
			curDir = nextDir[strconv.Itoa(curDir[0])+","+strconv.Itoa(curDir[1])]
			nr, nc = curRow+curDir[0], curCol+curDir[1]
		}

		visited[strconv.Itoa(nr)+","+strconv.Itoa(nc)] = true
		curRow, curCol = nr, nc
		nr, nc = curRow+curDir[0], curCol+curDir[1]
	}

	return count
}

func partTwo(grid [][]string) int {
	defer timeTrack(time.Now(), "partTwo")

	//Get start params
	curRow, curCol := 0, 0
	curDir := []int{0, 0}
	for r, row := range grid {
		for c, char := range row {
			if val, ok := startChars[char]; ok {
				curRow, curCol = r, c
				curDir = val
			}
		}
	}

	visited, _ := getPath(grid, curRow, curCol, curDir, false)
	blockers := make(map[string]bool)
	for k, _ := range visited {
		fields := strings.Split(k, ",")
		curRow, _ = strconv.Atoi(fields[0])
		curCol, _ = strconv.Atoi(fields[1])
		xDir, _ := strconv.Atoi(fields[2])
		yDir, _ := strconv.Atoi(fields[3])
		curDir = nextDir[strconv.Itoa(xDir)+","+strconv.Itoa(yDir)]

		gridCopy = resetGridCopy(grid)
		if curRow+xDir > -1 && curRow+xDir < len(grid) && curCol+yDir > -1 && curCol+yDir < len(grid[0]) {
			gridCopy[curRow+xDir][curCol+yDir] = "#"
			_, cycle := getPath(gridCopy, curRow, curCol, curDir, true)
			if cycle {
				blockers[strconv.Itoa(curRow+xDir)+","+strconv.Itoa(curCol+yDir)] = true
			}
		}
	}

	return len(blockers)
}

func resetGridCopy(grid [][]string) [][]string {
	gridCopy = make([][]string, len(grid))
	for i, row := range grid {
		gridCopy[i] = make([]string, len(row))
		for j, val := range row {
			gridCopy[i][j] = val
		}
	}
	return gridCopy
}

func getPath(grid [][]string, curRow int, curCol int, curDir []int, cycleCheck bool) (map[string]bool, bool) {
	visited := make(map[string]bool)
	visited[strconv.Itoa(curRow)+","+strconv.Itoa(curCol)+","+strconv.Itoa(curDir[0])+","+strconv.Itoa(curDir[1])] = true
	nr, nc := curRow+curDir[0], curCol+curDir[1]
	for nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[0]) {
		for grid[nr][nc] == "#" {
			curDir = nextDir[strconv.Itoa(curDir[0])+","+strconv.Itoa(curDir[1])]
			nr, nc = curRow+curDir[0], curCol+curDir[1]
		}

		if _, ok := visited[strconv.Itoa(nr)+","+strconv.Itoa(nc)+","+strconv.Itoa(curDir[0])+","+strconv.Itoa(curDir[1])]; cycleCheck && ok {
			return visited, true
		}

		visited[strconv.Itoa(nr)+","+strconv.Itoa(nc)+","+strconv.Itoa(curDir[0])+","+strconv.Itoa(curDir[1])] = true
		curRow, curCol = nr, nc
		nr, nc = curRow+curDir[0], curCol+curDir[1]
	}

	return visited, false
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
