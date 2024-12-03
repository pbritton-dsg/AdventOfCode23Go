package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i, maxCard, totalCards := 1, 1, 0
	myMap := make(map[int]int, 0)
	for scanner.Scan() && i <= maxCard+1 {
		line := scanner.Text()
		line = line[strings.Index(line, ":")+2:]
		lists := strings.Split(line, "|")
		lists[0] = strings.TrimSpace(lists[0])
		lists[1] = strings.TrimSpace(lists[1])
		winningNums := make([]int, 0)

		_, ok := myMap[i]
		if !ok {
			myMap[i] = 1
		} else {
			myMap[i] += 1
		}

		for _, winningNum := range strings.Split(lists[0], " ") {
			n, _ := strconv.Atoi(winningNum)
			winningNums = append(winningNums, n)
		}

		matches := 0
		for _, myNum := range strings.Split(lists[1], " ") {
			n, err := strconv.Atoi(myNum)
			if err == nil && slices.Contains(winningNums, n) {
				matches += 1
			}
		}

		counter := matches
		for counter > 0 {
			_, ok := myMap[i+counter]
			if !ok {
				myMap[i+counter] = myMap[i]
			} else {
				myMap[i+counter] += myMap[i]
			}
			counter -= 1
		}
		maxCard = max(maxCard, i+matches)
		i += 1
	}

	for _, val := range myMap {
		totalCards += val
	}
	fmt.Println(totalCards)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
