package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	antennaeMap := make(map[rune][][]int)
	grid := make([][]rune, 0)
	rowCount := 0
	for scanner.Scan() {
		input := scanner.Text()
		row := make([]rune, len(input))
		for col, val := range input {
			row[col] = val
			if val != '.' {
				antennaeMap[val] = append(antennaeMap[val], []int{rowCount, col})
			}
		}
		grid = append(grid, row)
		rowCount++
	}

	fmt.Println(partOne(grid, antennaeMap))
	fmt.Println(partTwo(grid, antennaeMap))
}

func partOne(grid [][]rune, antennaeMap map[rune][][]int) int {
	defer timeTrack(time.Now(), "partOne")

	resultsMap := make(map[string]bool)
	for _, positions := range antennaeMap {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				antiNodes := checkAntiNode(positions[i], positions[j])
				for _, node := range antiNodes {
					if node[0] > -1 && node[0] < len(grid) && node[1] > -1 && node[1] < len(grid[0]) {
						resultsMap[strconv.Itoa(node[0])+","+strconv.Itoa(node[1])] = true
					}
				}
			}
		}
	}

	return len(resultsMap)
}

func partTwo(grid [][]rune, antennaeMap map[rune][][]int) int {
	defer timeTrack(time.Now(), "partOne")

	resultsMap := make(map[string]bool)
	for _, positions := range antennaeMap {
		for i := 0; i < len(positions); i++ {
			resultsMap[strconv.Itoa(positions[i][0])+","+strconv.Itoa(positions[i][1])] = true
			for j := i + 1; j < len(positions); j++ {
				antiNodes := checkAntiNodePt2(positions[i], positions[j], grid)
				for _, node := range antiNodes {
					resultsMap[strconv.Itoa(node[0])+","+strconv.Itoa(node[1])] = true
				}
			}
		}
	}

	return len(resultsMap)
}

func checkAntiNodePt2(p1, p2 []int, grid [][]rune) [][]int {
	xMin, xMax := min(p1[0], p2[0]), max(p1[0], p2[0])
	yMin, yMax := min(p1[1], p2[1]), max(p1[1], p2[1])

	xDif := int(math.Abs(float64(p1[0] - p2[0])))
	yDif := int(math.Abs(float64(p1[1] - p2[1])))

	antiNodes := make([][]int, 0)
	if Eq(p1, []int{xMin, yMin}) || Eq(p1, []int{xMax, yMax}) {
		for xMin-xDif > -1 && xMin-xDif < len(grid) && yMin-yDif > -1 && yMin-yDif < len(grid[0]) {
			antiNodes = append(antiNodes, []int{xMin - xDif, yMin - yDif})
			xMin = xMin - xDif
			yMin = yMin - yDif
		}
		for xMax+xDif > -1 && xMax+xDif < len(grid) && yMax+yDif > -1 && yMax+yDif < len(grid[0]) {
			antiNodes = append(antiNodes, []int{xMax + xDif, yMax + yDif})
			xMax = xMax + xDif
			yMax = yMax + yDif
		}

	} else {
		for xMin-xDif > -1 && xMin-xDif < len(grid) && yMax+yDif > -1 && yMax+yDif < len(grid[0]) {
			antiNodes = append(antiNodes, []int{xMin - xDif, yMax + yDif})
			xMin = xMin - xDif
			yMax = yMax + yDif

		}
		for xMax+xDif > -1 && xMax+xDif < len(grid) && yMin-yDif > -1 && yMin-yDif < len(grid[0]) {
			antiNodes = append(antiNodes, []int{xMax + xDif, yMin - yDif})
			xMax = xMax + xDif
			yMin = yMin - yDif
		}
	}
	return antiNodes
}

func checkAntiNode(p1, p2 []int) [][]int {
	xMin, xMax := min(p1[0], p2[0]), max(p1[0], p2[0])
	yMin, yMax := min(p1[1], p2[1]), max(p1[1], p2[1])

	xDif := int(math.Abs(float64(p1[0] - p2[0])))
	yDif := int(math.Abs(float64(p1[1] - p2[1])))

	if Eq(p1, []int{xMin, yMin}) || Eq(p1, []int{xMax, yMax}) {
		return [][]int{{xMin - xDif, yMin - yDif}, {xMax + xDif, yMax + yDif}}
	} else {
		return [][]int{{xMin - xDif, yMax + yDif}, {xMax + xDif, yMin - yDif}}
	}
}

func Eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
