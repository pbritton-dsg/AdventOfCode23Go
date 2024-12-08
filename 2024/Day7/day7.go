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

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	results = make(map[int]bool)
	equations := make(map[int][]int)
	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), ":")
		fields := strings.Fields(sides[1])
		result, _ := strconv.Atoi(sides[0])
		numbers := make([]int, len(fields))
		for i := range fields {
			n, _ := strconv.Atoi(fields[i])
			numbers[i] = n
		}
		equations[result] = numbers
	}

	fmt.Println(partOne(equations))
	//fmt.Println(partTwo(grid))
}

func partOne(equations map[int][]int) int {
	defer timeTrack(time.Now(), "partOne")

	sum := 0
	for result, values := range equations {
		sum += checkResult(result, values)
	}

	return sum
}

func checkResult(result int, values []int) int {
	if len(values) == 1 {
		if values[0] == result {
			return result
		} else {
			return 0
		}
	}

	n0, n1 := values[0], values[1]
	n3, _ := strconv.Atoi(strconv.Itoa(n0) + strconv.Itoa(n1))
	values = values[2:]

	y := append([]int{n0 + n1}, values...)
	z := append([]int{n0 * n1}, values...)
	w := append([]int{n3}, values...)

	return max(checkResult(result, y), checkResult(result, z), checkResult(result, w))

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
