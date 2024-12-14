package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var visited [][]uint8

var neighbors = [][]int{
	{-1, 0}, {0, -1}, {0, 1}, {1, 0},
}

type Point struct {
	x int
	y int
}

var currPlot []Point

func main() {
	f, _ := os.Open("inputs/testInput.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	garden := make([][]rune, 0)
	for scanner.Scan() {
		row := make([]rune, len(scanner.Text()))
		for i, c := range scanner.Text() {
			row[i] = c
		}
		garden = append(garden, row)
	}

	//fmt.Println(partOne(garden))
	fmt.Println(partTwo(garden))
}

func partOne(garden [][]rune) int {
	defer timeTrack(time.Now(), "partOne")
	resetVisited(garden)

	sum := 0

	plots := make([][]Point, 0)
	for row, _ := range garden {
		for col, _ := range garden[row] {
			currPlot = make([]Point, 0)
			getPlot(row, col, garden)
			if len(currPlot) > 0 {
				plots = append(plots, currPlot)
			}
		}
	}

	for _, plot := range plots {
		area := len(plot)
		perimeter := 0
		for _, point := range plot {
			perimeter += 4
			for _, neighbor := range neighbors {
				nr, nc := point.x+neighbor[0], point.y+neighbor[1]
				if nr > -1 && nr < len(garden) && nc > -1 && nc < len(garden[0]) {
					if garden[nr][nc] == garden[point.x][point.y] {
						perimeter--
					}
				}
			}
		}
		fmt.Println("Area: ", area, "Perimeter: ", perimeter)
		sum += perimeter * area

	}

	return sum
}

func partTwo(garden [][]rune) int {
	defer timeTrack(time.Now(), "partOne")
	resetVisited(garden)

	sum := 0

	plots := make([][]Point, 0)
	for row, _ := range garden {
		for col, _ := range garden[row] {
			currPlot = make([]Point, 0)
			getPlot(row, col, garden)
			if len(currPlot) > 0 {
				plots = append(plots, currPlot)
			}
		}
	}

	for _, plot := range plots {
		area := len(plot)
		sides := 0
		minX, minY, maxX, maxY := 999, 999, -999, -999
		for _, point := range plot {
			minX = min(minX, point.x)
			minY = min(minY, point.y)
			maxX = max(maxX, point.x)
			maxY = max(maxY, point.y)
			//for _, neighbor := range neighbors {
			//	nr, nc := point.x+neighbor[0], point.y+neighbor[1]
			//	if nr > -1 && nr < len(garden) && nc > -1 && nc < len(garden[0]) {
			//		if garden[nr][nc] == garden[point.x][point.y] {
			//			sides--
			//		}
			//	}
			//}
		}
		fmt.Println("Area: ", area, "Dx: ", maxX-minX, "Dy: ", maxY-minY)
		sum += sides * area

	}

	return sum
}

func getPlot(row int, col int, garden [][]rune) {
	if visited[row][col] != 1 {
		currPlot = append(currPlot, Point{row, col})
	}
	visited[row][col] = 1
	for _, neighbor := range neighbors {
		nr, nc := row+neighbor[0], col+neighbor[1]
		if nr > -1 && nr < len(garden) && nc > -1 && nc < len(garden[row]) && visited[nr][nc] != 1 && garden[nr][nc] == garden[row][col] {
			getPlot(nr, nc, garden)
		}
	}
}

func resetVisited(grid [][]rune) {
	visited = make([][]uint8, len(grid))
	for i := range grid {
		visited[i] = make([]uint8, len(grid[i]))
	}
}

//func countTotal(steps []Step) int {
//	defer timeTrack(time.Now(), "countTotal")
//	x, y, perimeter, area := 0, 0, 0, 0
//	for _, step := range steps {
//		dy, dx := curDirPt2[step.dir][0], curDirPt2[step.dir][1]
//		dy, dx = dy*step.length, dx*step.length
//		x, y = x+dx, y+dy
//		perimeter += step.length
//		area += x * dy
//	}
//
//	return area + perimeter/2 + 1
//}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
