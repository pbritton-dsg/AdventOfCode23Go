package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fileSystem := make([]int, 0)
	for scanner.Scan() {
		input := scanner.Text()
		for _, c := range input {
			num, _ := strconv.Atoi(string(c))
			fileSystem = append(fileSystem, num)
		}
	}

	fmt.Println(partOne(fileSystem))
	fmt.Println(partTwo(fileSystem))
}

func partOne(filesystem []int) int {
	defer timeTrack(time.Now(), "partOne")
	sum := 0

	endFileVal := len(filesystem) / 2
	startFileVal := 0
	leftIdx, rightIdx := 0, len(filesystem)-1

	overallIdx, filesToMove, openSlots := 0, 0, 0
	for leftIdx < rightIdx {
		if leftIdx%2 == 0 {
			fileCount := filesystem[leftIdx]
			for fileCount > 0 {
				sum += overallIdx * startFileVal
				fileCount--
				overallIdx++
			}
			leftIdx++
			startFileVal++
		} else {
			if filesToMove == 0 {
				filesToMove = filesystem[rightIdx]
			}
			openSlots = filesystem[leftIdx]
			for openSlots > 0 {
				sum += overallIdx * endFileVal

				filesToMove--
				openSlots--
				overallIdx++

				if filesToMove == 0 {
					rightIdx = rightIdx - 2
					filesToMove = filesystem[rightIdx]
					endFileVal--
				}
			}
			leftIdx++
		}
	}

	for filesToMove > 0 {
		sum += overallIdx * endFileVal
		overallIdx++
		filesToMove--
	}

	return sum
}

func partTwo(filesystem []int) int {
	defer timeTrack(time.Now(), "partTwo")

	remainingSlots := make([]int, len(filesystem))
	for i, val := range filesystem {
		remainingSlots[i] = val
	}
	sum := 0

	endFileVal := len(filesystem) / 2
	rightIdx := len(filesystem) - 1

	for rightIdx > 0 {

		position := 0
		for i := 0; i < rightIdx; i++ {
			position += filesystem[i]
		}

		filesToMove := filesystem[rightIdx]
		filesMoved := false
		openFileIdx := 1
		for openFileIdx < rightIdx {
			openSlots := remainingSlots[openFileIdx]
			if filesToMove <= openSlots {

				movedPosition := 0
				for i := 0; i < openFileIdx; i++ {
					movedPosition += filesystem[i]
				}

				movedPosition += filesystem[openFileIdx] - remainingSlots[openFileIdx]

				for filesToMove > 0 {
					sum += movedPosition * endFileVal
					movedPosition++
					filesToMove--
					openSlots--
				}
				remainingSlots[openFileIdx] = openSlots - filesToMove
				filesMoved = true
				break
			}
			openFileIdx += 2
		}

		if !filesMoved {
			for filesToMove > 0 {
				sum += position * endFileVal
				position++
				filesToMove--
			}
		}

		rightIdx -= 2
		endFileVal--
	}

	return sum
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
