package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	fields := make([]string, 0)
	for scanner.Scan() {
		fields = append(fields, strings.Split(scanner.Text(), "mul(")...)
	}

	fmt.Println(partOne(fields))
	fmt.Println(partTwo(fields))
}

func partOne(fields []string) int {
	defer timeTrack(time.Now(), "partOne")
	sum := 0
	var validMul = regexp.MustCompile(`^\d{1,3},\d{1,3}\).*$`)
	for _, field := range fields {
		if validMul.MatchString(field) {
			digits := strings.Split(field[:strings.IndexByte(field, ')')], ",")
			first, _ := strconv.Atoi(digits[0])
			second, _ := strconv.Atoi(digits[1])
			sum += first * second
		}
	}
	return sum
}

func partTwo(fields []string) int {
	defer timeTrack(time.Now(), "partOne")
	sum := 0
	enabled := true
	var validMul = regexp.MustCompile(`^\d{1,3},\d{1,3}\).*$`)
	for _, field := range fields {
		if validMul.MatchString(field) {
			if enabled {
				digits := strings.Split(field[:strings.IndexByte(field, ')')], ",")
				first, _ := strconv.Atoi(digits[0])
				second, _ := strconv.Atoi(digits[1])
				sum += first * second
			}
		}

		lastDo := strings.LastIndex(field, "do()")
		lastDont := strings.LastIndex(field, "don't()")
		if lastDo > 0 || lastDont > 0 {
			enabled = lastDo > lastDont
		}

	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
