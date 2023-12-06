package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main2() {
	f, err := os.Open("inputs/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum := 0
		for _, char := range line {
			if unicode.IsNumber(char) {
				n, _ := strconv.Atoi(string(char))
				lineNum += n * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				n, _ := strconv.Atoi(string(line[i]))
				lineNum += n
				break
			}
		}

		sum += lineNum

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(sum)
}
