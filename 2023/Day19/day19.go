package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	x int
	m int
	a int
	s int
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	part1 := true
	rules := make(map[string][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			part1 = false
			continue
		}
		if part1 {
			rules[line[0:strings.Index(line, "{")]] = strings.Split(line[strings.Index(line, "{")+1:len(line)-1], ",")
		} else {
			steps := strings.Split(line[1:len(line)-1], ",")
			stepsMap := make(map[string]int, 0)
			for i := 0; i < 4; i++ {
				num, _ := strconv.Atoi(steps[i][strings.Index(steps[i], "=")+1 : len(steps[i])])
				stepsMap[steps[i][0:1]] = num
			}
			sum += countTotal(stepsMap, rules)
		}
	}
	fmt.Println(sum)
}

func countTotal(step map[string]int, rules map[string][]string) int {
	current := rules["in"]
	complete := -1
	i := 0
	for i < len(current) && complete < 0 {
		rule := current[i]
		_, ok := rules[rule]
		if ok || rule == "A" || rule == "R" {
			current, complete = getNextRule(rule, rules, step)
			i = 0
			continue
		}
		if evalRule(step, rule) {
			current, complete = getNextRule(rule[strings.Index(rule, ":")+1:], rules, step)
		} else {
			i += 1
		}
	}
	return complete
}

func getNextRule(key string, rules map[string][]string, step map[string]int) ([]string, int) {
	if key == "A" {
		return []string{}, step["x"] + step["m"] + step["a"] + step["s"]
	} else if key == "R" {
		return []string{}, 0
	}
	return rules[key], -1
}

func evalRule(step map[string]int, rule string) bool {
	varToEval := rule[0:1]
	operator := rule[1:2]
	num, _ := strconv.Atoi(rule[2:strings.Index(rule, ":")])
	if operator == "<" {
		return step[varToEval] < num
	} else {
		return step[varToEval] > num
	}
	return true
}
