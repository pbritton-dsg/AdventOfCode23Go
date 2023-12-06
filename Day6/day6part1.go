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

	i := 0
	times, records := make([]int, 0), make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = line[strings.Index(line, ":")+1:]
		fields := strings.Fields(line)

		for _, val := range fields {
			n, _ := strconv.Atoi(val)
			if i == 0 {
				times = append(times, n)
			} else {
				records = append(records, n)
			}
		}
		i += 1
	}

	product := 1
	for j := 0; j < len(times); j++ {
		product *= timeRange1(times[j], records[j])
	}
	fmt.Println(product)
}

func timeRange1(time int, record int) int {
	l, r := 0, time
	leftWin, rightWin := false, false
	for !leftWin && !rightWin {
		leftWin = l*(time-l) > record
		rightWin = r*(time-r) > record
		if !leftWin {
			l += 1
		}
		if !rightWin {
			r -= 1
		}
	}
	fmt.Println(l, r)
	return r - l + 1
}
