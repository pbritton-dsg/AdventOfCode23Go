package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main1() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	possibleGames := make([]int, 0)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		line = line[strings.Index(line, ":")+2:]
		pulls := strings.Split(line, ";")
		var maxCubes = map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		}
		isPossible := true
		for _, pull := range pulls {
			combos := strings.Split(pull, ",")
			for _, val := range combos {
				val := strings.TrimSpace(val)
				combo := strings.Split(val, " ")
				n, _ := strconv.Atoi(combo[0])
				if maxCubes[combo[1]] < n {
					isPossible = false
				}
			}
		}

		if isPossible {
			possibleGames = append(possibleGames, i)
		}
		i += 1
	}

	sum := 0
	for _, val := range possibleGames {
		sum += val
	}
	fmt.Println(sum)
}
