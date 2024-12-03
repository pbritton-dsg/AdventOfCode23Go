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

var curDir = map[string][]int{
	"R": {0, 1},
	"L": {0, -1},
	"D": {1, 0},
	"U": {-1, 0},
}

var curDirPt2 = map[string][]int{
	"0": {0, 1},
	"2": {0, -1},
	"1": {1, 0},
	"3": {-1, 0},
}

type Step struct {
	dir    string
	length int
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	part1 := false
	steps := make([]Step, 0)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		dir := ""
		dist := 0
		if part1 {
			dir = fields[0]
			dist, _ = strconv.Atoi(fields[1])
		} else {
			dir = fields[2][len(fields[2])-2 : len(fields[2])-1]
			n, _ := strconv.ParseInt(fields[2][2:7], 16, 64)
			dist = int(n)
		}
		steps = append(steps, Step{dir, dist})
	}
	fmt.Println(countTotal(steps))
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func countTotal(steps []Step) int {
	defer timeTrack(time.Now(), "countTotal")
	x, y, perimeter, area := 0, 0, 0, 0
	for _, step := range steps {
		dy, dx := curDirPt2[step.dir][0], curDirPt2[step.dir][1]
		dy, dx = dy*step.length, dx*step.length
		x, y = x+dx, y+dy
		perimeter += step.length
		area += x * dy
	}

	return area + perimeter/2 + 1
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
