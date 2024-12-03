package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	fields := make([]string, 0)
	for scanner.Scan() {
		fields = strings.Split(scanner.Text(), ",")
	}

	fmt.Println(partOne(fields))
	fmt.Println(partTwo(fields))
}

func partOne(fields []string) int {
	result := 0
	for _, field := range fields {
		total := 0
		for _, c := range field {
			total += int(c)
			total *= 17
			total %= 256
		}
		result += total
	}

	return result
}

func partTwo(fields []string) int {
	myMap := make(map[int][]string)
	total := 0
	for _, field := range fields {
		box := 0
		remove := strings.Contains(field, "-")
		for _, c := range field {
			if string(c) == "-" || string(c) == "=" {
				break
			}
			box += int(c)
			box *= 17
			box %= 256
		}
		contents := ""
		if remove {
			contents = strings.Replace(field, "-", " ", 1)
		} else {
			contents = strings.Replace(field, "=", " ", 1)
		}
		label := strings.Split(contents, " ")[0]

		if val, ok := myMap[box]; !ok && !remove {
			myMap[box] = []string{contents}
		} else {
			index := slices.IndexFunc(val, func(s string) bool {
				return strings.Split(s, " ")[0] == label
			})
			if !remove {
				if index != -1 {
					myMap[box][index] = contents
				} else {
					myMap[box] = append(val, contents)
				}
			} else {
				myMap[box] = slices.DeleteFunc(val, func(s string) bool {
					return strings.Split(s, " ")[0] == label
				})
			}
		}
	}

	keys := make([]int, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		for i, n := range myMap[k] {
			length, _ := strconv.Atoi(strings.Split(n, " ")[1])
			total += (k + 1) * (i + 1) * length
		}
	}

	return total
}
