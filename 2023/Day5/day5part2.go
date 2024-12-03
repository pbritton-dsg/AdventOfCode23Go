package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	seeds := getSeedList(line)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, ":") > -1 {
			mapLines := make([]string, 0)
			for scanner.Scan() && scanner.Text() != "" {
				mapLines = append(mapLines, scanner.Text())
			}
			seeds = mapToDest(seeds, mapLines)
		}
	}

	minSeed := math.MaxInt
	for _, seed := range seeds {
		if seed[0] < minSeed {
			minSeed = seed[0]
		}
	}
	fmt.Println(minSeed)
}

func mapToDest(seeds [][]int, mapLines []string) [][]int {
	rangeMap := buildMap(mapLines)
	newNums := make([][]int, 0)
	for len(seeds) > 0 {
		s, e := seeds[0][0], seeds[0][1]
		seeds = seeds[1:]
		overlap := false
		for _, r := range rangeMap {
			destStart, sourceStart, size := r[0], r[1], r[2]
			os := max(s, sourceStart)
			oe := min(e, sourceStart+size)
			if os < oe {
				overlap = true
				newNums = append(newNums, []int{os - sourceStart + destStart, oe - sourceStart + destStart})
				if os > s {
					seeds = append(seeds, []int{s, os})
				}
				if e > oe {
					seeds = append(seeds, []int{oe, e})
				}
				break
			}
		}
		if !overlap {
			newNums = append(newNums, []int{s, e})
		}
	}

	return newNums
}

func buildMap(mapLines []string) [][]int {
	myMap := make([][]int, 0)
	for _, line := range mapLines {
		nums := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(nums[0])
		sourceStart, _ := strconv.Atoi(nums[1])
		size, _ := strconv.Atoi(nums[2])
		myMap = append(myMap, []int{destStart, sourceStart, size})
	}
	return myMap
}

func getSeedList(line string) [][]int {
	seeds := make([][]int, 0)
	line = line[strings.Index(line, ":")+2:]
	arr := strings.Split(strings.TrimSpace(line), " ")
	for len(arr) > 0 {
		start, _ := strconv.Atoi(arr[0])
		l, _ := strconv.Atoi(arr[1])
		seeds = append(seeds, []int{start, start + l})
		arr = arr[2:]
	}
	return seeds
}
