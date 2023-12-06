package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main1() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	currentNums := getSeedList1(line)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, ":") > -1 {
			mapLines := make([]string, 0)
			for scanner.Scan() && scanner.Text() != "" {
				mapLines = append(mapLines, scanner.Text())
			}
			currentNums = getNewNums1(currentNums, mapLines)
		}
	}

	minNum := math.MaxInt
	for _, num := range currentNums {
		if num < minNum {
			minNum = num
		}
	}

	fmt.Println(minNum)
}

func getNewNums1(currentNums []int, mapLines []string) []int {
	myMap := make([][]int, 0)
	for _, line := range mapLines {
		nums := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(nums[0])
		sourceStart, _ := strconv.Atoi(nums[1])
		size, _ := strconv.Atoi(nums[2])
		myMap = append(myMap, []int{destStart, sourceStart, size})
	}

	newNums := make([]int, 0)
	for _, num := range currentNums {
		for _, mapLine := range myMap {
			if num >= mapLine[1] && num <= mapLine[1]+mapLine[2]-1 {
				num = mapLine[0] + num - mapLine[1]
				break
			}
		}
		newNums = append(newNums, num)
	}

	return newNums
}

func getSeedList1(line string) []int {
	seeds := make([]int, 0)
	line = line[strings.Index(line, ":")+2:]
	line = strings.TrimSpace(line)
	for _, seed := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(seed)
		seeds = append(seeds, n)
	}
	return seeds
}
