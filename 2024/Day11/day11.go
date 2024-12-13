package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	stones := make(map[int]int, 0)
	for scanner.Scan() {
		charNums := strings.Fields(scanner.Text())
		for _, char := range charNums {
			num, _ := strconv.Atoi(char)
			if _, ok := stones[num]; ok {
				stones[num]++
			} else {
				stones[num] = 1
			}
		}
	}

	//fmt.Println(partOne(stones, 75))
	fmt.Println(partTwo(stones, 75))
}

func partOne(stones []int, blinks int) int {
	defer timeTrack(time.Now(), "partOne")

	for blinks > 0 {
		i := 0
		for i < len(stones) {
			stoneString := strconv.Itoa(stones[i])
			if stones[i] == 0 {
				stones[i] = 1
			} else if len(stoneString)%2 == 0 {
				mid := len(stoneString) / 2
				s1, _ := strconv.Atoi(stoneString[:mid])
				s2, _ := strconv.Atoi(stoneString[mid:])
				newStones := []int{s1, s2}
				stones = append(stones[:i], append(newStones, stones[i+1:]...)...)
				i++
			} else {
				stones[i] = stones[i] * 2024
			}
			i++
		}
		blinks--
	}

	return len(stones)
}

func partTwo(stones map[int]int, blinks int) int {
	defer timeTrack(time.Now(), "partTwo")

	for blinks > 0 {
		stoneCopy := maps.Clone(stones)
		for stoneNum, count := range stoneCopy {
			stoneString := strconv.Itoa(stoneNum)
			if stoneNum == 0 {
				stones[1] += count
				stones[0] -= count
			} else if len(stoneString)%2 == 0 {
				mid := len(stoneString) / 2
				s1, _ := strconv.Atoi(stoneString[:mid])
				s2, _ := strconv.Atoi(stoneString[mid:])
				stones[s1] += count
				stones[s2] += count
				stones[stoneNum] -= count
			} else {
				stones[2024*stoneNum] += count
				stones[stoneNum] -= count
			}
		}
		blinks--
	}

	stoneCount := 0
	for _, count := range stones {
		stoneCount += count
	}
	return stoneCount
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
