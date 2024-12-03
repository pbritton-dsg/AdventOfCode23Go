package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	part2 := true
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		springs, arrangements := input[0], input[1]
		if part2 {
			s, a := "", ""
			for i := 0; i < 5; i++ {
				s, a = s+springs+"?", a+arrangements+","
			}
			springs, arrangements = strings.TrimSuffix(s, "?"), strings.TrimSuffix(a, ",")
		}

		blocks := strings.Split(arrangements, ",")
		springsArr := make([]string, len(springs))
		for i, r := range springs {
			springsArr[i] = string(r)
		}
		possible := make([]int, len(blocks))
		for i, a := range blocks {
			if n, err := strconv.Atoi(a); err == nil {
				possible[i] = n
			}
		}

		dp := make(map[string]int)
		total += count(springsArr, possible, 0, 0, 0, dp)
		fmt.Println(total)
	}
	fmt.Println(total)
}

func count(springs []string, blocks []int, i, bi, current int, dp map[string]int) int {
	key := strconv.Itoa(i) + "," + strconv.Itoa(bi) + "," + strconv.Itoa(current)
	if val, ok := dp[key]; ok {
		return val
	}

	if i == len(springs) {
		if bi == len(blocks) && current == 0 {
			return 1
		} else if bi == len(blocks)-1 && blocks[bi] == current {
			return 1
		} else {
			return 0
		}
	}
	ans := 0
	for _, spot := range []string{".", "#"} {
		if springs[i] == spot || springs[i] == "?" {
			if spot == "." && current == 0 {
				ans += count(springs, blocks, i+1, bi, 0, dp)
			} else if spot == "." && current > 0 && bi < len(blocks) && blocks[bi] == current {
				ans += count(springs, blocks, i+1, bi+1, 0, dp)
			} else if spot == "#" {
				ans += count(springs, blocks, i+1, bi, current+1, dp)
			}
		}
	}
	dp[key] = ans
	return ans
}
