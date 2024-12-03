package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 0
	time, record := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		line = line[strings.Index(line, ":")+1:]
		num := strings.Join(strings.Fields(line), "")
		if i == 0 {
			time, _ = strconv.Atoi(num)
		} else {
			record, _ = strconv.Atoi(num)
		}
		i += 1
	}

	fmt.Println(timeRange(time, record))
}

func timeRange(time int, record int) int {
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
	return r - l + 1
}
