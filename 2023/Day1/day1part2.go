package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var numMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	f, _ := os.Open("inputs/part2Input.txt")

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum := 0

		firstIndex, firstStringNum := math.MaxInt, -1
		for key, val := range numMap {
			index := strings.Index(line, key)
			if index >= 0 && index < firstIndex {
				firstStringNum = val
				firstIndex = strings.Index(line, key)
			}
		}

		firstNumIndex, firstNum := math.MaxInt, -1
		for pos, char := range line {
			if unicode.IsNumber(char) {
				n, _ := strconv.Atoi(string(char))
				firstNum = n
				firstNumIndex = pos
				break
			}
		}

		if firstNumIndex > firstIndex {
			lineNum += firstStringNum * 10
		} else {
			lineNum += firstNum * 10
		}

		lastIndex, lastStringNum := -1, -1
		for key, val := range numMap {
			if strings.LastIndex(line, key) > lastIndex {
				lastStringNum = val
				lastIndex = strings.LastIndex(line, key)
			}
		}

		lastNumIndex, lastNum := -1, -1
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				n, _ := strconv.Atoi(string(line[i]))
				lastNum = n
				lastNumIndex = i
				break
			}
		}

		if lastNumIndex < lastIndex {
			lineNum += lastStringNum
		} else {
			lineNum += lastNum
		}

		sum += lineNum
	}
	fmt.Println(sum)
}
