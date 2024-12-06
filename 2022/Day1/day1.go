package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	maxElf, elfSum := 0, 0
	for scanner.Scan() {
		curAmt, err := strconv.Atoi(scanner.Text())
		if err == nil {
			elfSum += curAmt
		} else {
			if elfSum > maxElf {
				maxElf = elfSum
			}
			elfSum = 0
		}
	}

	fmt.Println(maxElf)

	//fmt.Println(partOne(grid))
	//fmt.Println(partTwo(grid))
}
