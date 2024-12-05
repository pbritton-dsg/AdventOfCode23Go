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

	pageMap := make(map[int][]int)

	rules := make([][]int, 0)
	buildingMap := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			buildingMap = false
			continue
		}
		if buildingMap {
			fields := strings.Split(line, "|")
			x, _ := strconv.Atoi(fields[0])
			y, _ := strconv.Atoi(fields[1])

			if val, ok := pageMap[x]; ok {
				pageMap[x] = append(val, y)
			} else {
				pageMap[x] = []int{y}
			}
		} else {
			fields := strings.Split(line, ",")
			rule := make([]int, 0)
			for _, field := range fields {
				x, _ := strconv.Atoi(field)
				rule = append(rule, x)
			}
			rules = append(rules, rule)
		}
	}

	//fmt.Println(partOne(rules, pageMap))
	fmt.Println(partTwo(rules, pageMap))
}

func partOne(rules [][]int, pageMap map[int][]int) int {
	defer timeTrack(time.Now(), "partOne")
	sum := 0
	for _, rule := range rules {
		validRule := true
		for i := 1; i < len(rule); i++ {
			val, ok := pageMap[rule[i]]
			for j := i - 1; j >= 0; j-- {
				if ok && slices.Contains(val, rule[j]) {
					fmt.Println(rule)
					validRule = false
					break
				}
			}
			if !validRule {
				break
			}
		}
		if validRule {
			sum += rule[len(rule)/2]
		}
	}

	return sum
}

func partTwo(rules [][]int, pageMap map[int][]int) int {
	defer timeTrack(time.Now(), "partTwo")

	sum := 0
	for _, rule := range rules {
		//sum += fixRule(rule)
		validRule := true
		for i := 1; i < len(rule); i++ {
			val, ok := pageMap[rule[i]]
			for j := i - 1; j >= 0; j-- {
				if ok && slices.Contains(val, rule[j]) {
					sum += fixRule(rule, pageMap)
					validRule = false
					break
				}
			}
			if !validRule {
				break
			}
		}
	}

	return sum
}

func fixRule(rule []int, pageMap map[int][]int) int {

	currRuleMap := make(map[int][]int)
	for _, r := range rule {
		currRuleMap[r] = pageMap[r]
	}

	order := topSort(currRuleMap)
	newOrder := make([]int, 0)
	for _, x := range order {
		if slices.Contains(rule, x) {
			newOrder = append(newOrder, x)
		}
	}

	return newOrder[len(newOrder)/2]
}

//
//func fixRule(rule []int) int {
//	mapOrderCompare := func(a, b int) int {
//		return cmp.Compare(slices.Index(mapOrder, a), slices.Index(mapOrder, b))
//	}
//
//	val := 0
//
//	orderedRule := make([]int, len(rule))
//	copy(orderedRule, rule)
//
//	slices.SortFunc(orderedRule, mapOrderCompare)
//	//slices.SortFunc(rule, mapOrderCompare)
//
//	for i := 0; i < len(orderedRule); i++ {
//		if rule[i] != orderedRule[i] {
//			val += orderedRule[len(orderedRule)/2]
//			break
//		}
//	}
//
//	return val
//	//return rule[len(rule)/2]
//}

//func topSort(graph map[int][]int) []int {
//	inDegree := make(map[int]int)
//	for node := range graph {
//		inDegree[node] = 0
//	}
//
//	for _, neighbors := range graph {
//		for _, neighbor := range neighbors {
//			inDegree[neighbor]++
//		}
//	}
//
//	queue := list.New()
//	for node, degree := range inDegree {
//		if degree == 0 {
//			queue.PushBack(node)
//		}
//	}
//
//	result := []int{}
//	for queue.Len() > 0 {
//		node := queue.Remove(queue.Front()).(int)
//		result = append(result, node)
//
//		for _, neighbor := range graph[node] {
//			inDegree[neighbor]--
//			if inDegree[neighbor] == 0 {
//				queue.PushBack(neighbor)
//			}
//		}
//	}
//
//	return result
//}

func topSort(g map[int][]int) []int {
	linearOrder := []int{}
	inDegree := map[int]int{}

	for n := range g {
		inDegree[n] = 0
	}

	for _, adjacent := range g {
		for _, v := range adjacent {
			inDegree[v]++
		}
	}

	next := []int{}
	for u, v := range inDegree {
		if v != 0 {
			continue
		}

		next = append(next, u)
	}

	for len(next) > 0 {
		u := next[0]
		next = next[1:]

		linearOrder = append(linearOrder, u)

		for _, v := range g[u] {
			inDegree[v]--

			if inDegree[v] == 0 {
				next = append(next, v)
			}
		}
	}

	return linearOrder
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
