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
	{-1, 0}, {0, -1}, {0, 1}, {1, 0},
}

type Path struct {
	x, y  int
	route []string
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		line := scanner.Text()
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	fmt.Println(partOne(grid))
	//fmt.Println(partTwo(rules, pageMap))
}

func partOne(grid [][]int) int {
	defer timeTrack(time.Now(), "partOne")

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				sum += len(findPath(grid, row, col))
			}
		}
	}

	return sum
}

func findPath(grid [][]int, row, col int) [][]string {
	position := strconv.Itoa(row) + "," + strconv.Itoa(col)
	queue := []Path{{x: row, y: col, route: []string{position}}}
	//allPaths := make(map[string]bool)
	var allPaths [][]string

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		position = strconv.Itoa(curr.x) + "," + strconv.Itoa(curr.y)
		if grid[curr.x][curr.y] == 9 {
			//allPaths[position] = true
			allPaths = append(allPaths, curr.route)
			continue
		}

		for _, neighbor := range neighbors {
			nr, nc := curr.x+neighbor[0], curr.y+neighbor[1]
			position = strconv.Itoa(nr) + "," + strconv.Itoa(nc)
			if nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[row]) && grid[nr][nc] == grid[curr.x][curr.y]+1 {
				newRoute := append([]string{}, curr.route...)
				newRoute = append(newRoute, position)
				queue = append(queue, Path{x: nr, y: nc, route: newRoute})
			}
		}

	}

	return allPaths
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
