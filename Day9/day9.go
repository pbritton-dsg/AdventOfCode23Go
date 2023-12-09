package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   string
	left  string
	right string
}

type Graph struct {
	nodes map[string]Node
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sequences := make([][]int, 0)
	for scanner.Scan() {
		sequence := make([]int, 0)
		fields := strings.Fields(scanner.Text())
		for _, f := range fields {
			if n, err := strconv.Atoi(f); err == nil {
				sequence = append(sequence, n)
			}
		}
		sequences = append(sequences, sequence)
	}

	fmt.Println(partOne(sequences))
	fmt.Println(partTwo(sequences))
}

func partOne(sequences [][]int) int {
	sum := 0
	for _, line := range sequences {
		sum += findNext(line)
	}
	return sum
}

func partTwo(sequences [][]int) int {
	sum := 0
	for _, line := range sequences {
		sum += findPrevious(line)
	}
	return sum
}

func findNext(sequence []int) int {
	if allZeroes(sequence) {
		return 0
	}
	newSequence := make([]int, 0)
	for i := 1; i < len(sequence); i++ {
		newSequence = append(newSequence, sequence[i]-sequence[i-1])
	}
	return sequence[len(sequence)-1] + findNext(newSequence)
}

func findPrevious(sequence []int) int {
	if allZeroes(sequence) {
		return 0
	}
	newSequence := make([]int, 0)
	for i := 1; i < len(sequence); i++ {
		newSequence = append(newSequence, sequence[i]-sequence[i-1])
	}
	return sequence[0] - findPrevious(newSequence)
}

func allZeroes(sequence []int) bool {
	for _, s := range sequence {
		if s != 0 {
			return false
		}
	}
	return true
}
