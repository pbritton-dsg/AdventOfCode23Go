package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("inputs/part2Input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 1
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = line[strings.Index(line, ":")+2:]
		pulls := strings.Split(line, ";")
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, pull := range pulls {
			combos := strings.Split(pull, ",")
			for _, val := range combos {
				val := strings.TrimSpace(val)
				combo := strings.Split(val, " ")
				n, _ := strconv.Atoi(combo[0])
				minCubes[combo[1]] = max(minCubes[combo[1]], n)
			}
		}

		product := 1
		for _, val := range minCubes {
			product *= val
		}
		sum += product

		i += 1
	}

	fmt.Println(sum)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
