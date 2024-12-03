package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	reports := make([][]int, 0)
	for scanner.Scan() {
		report := make([]int, 0)
		fields := strings.Fields(scanner.Text())
		for i := range fields {
			num, _ := strconv.Atoi(fields[i])
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	fmt.Println(partOne(reports))
	fmt.Println(partTwo(reports))
}

func partOne(reports [][]int) int {
	defer timeTrack(time.Now(), "partOne")
	safeCount := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeCount++
		}
	}
	return safeCount
}

func isReportSafe(report []int) bool {
	direction := report[1] - report[0]
	for i := 1; i < len(report); i++ {
		dif := report[i] - report[i-1]
		if dif*direction <= 0 || dif < -3 || dif > 3 {
			return false
		}
	}
	return true
}

func partTwo(reports [][]int) int {
	defer timeTrack(time.Now(), "partTwo")
	safeCount := 0
	for _, report := range reports {
		for i := range report {
			if isReportSafe(slices.Delete(slices.Clone(report), i, i+1)) {
				safeCount++
				break
			}
		}
	}
	return safeCount
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
