package main

//
//import (
//	"bufio"
//	"cmp"
//	"fmt"
//	"github.com/dominikbraun/graph"
//	"log"
//	"os"
//	"slices"
//	"strconv"
//	"strings"
//	"time"
//)
//
//var mapOrder []int
//
//func main() {
//	f, _ := os.Open("inputs/input.txt")
//	defer f.Close()
//
//	scanner := bufio.NewScanner(f)
//
//	g := graph.New(graph.IntHash, graph.Directed(), graph.PreventCycles())
//	rules := make([][]int, 0)
//	buildingMap := true
//	vertexMap := make(map[int]int)
//	for scanner.Scan() {
//		line := scanner.Text()
//		if line == "" {
//			buildingMap = false
//			continue
//		}
//		if buildingMap {
//			fields := strings.Split(line, "|")
//			x, _ := strconv.Atoi(fields[0])
//			y, _ := strconv.Atoi(fields[1])
//
//			if _, ok := vertexMap[x]; !ok {
//				g.AddVertex(x)
//			}
//			if _, ok := vertexMap[y]; !ok {
//				g.AddVertex(y)
//			}
//
//			g.AddEdge(x, y)
//
//		} else {
//			fields := strings.Split(line, ",")
//			rule := make([]int, 0)
//			for _, field := range fields {
//				x, _ := strconv.Atoi(field)
//				rule = append(rule, x)
//			}
//			rules = append(rules, rule)
//		}
//	}
//
//	mapOrder, _ = graph.TopologicalSort(g)
//	fmt.Println(mapOrder)
//
//	//fmt.Println(partOne(rules, pageMap))
//	fmt.Println(partTwo(rules))
//}
//
//func partTwo(rules [][]int) int {
//	defer timeTrack(time.Now(), "partTwo")
//
//	sum := 0
//	for _, rule := range rules {
//		sum += fixRule(rule)
//	}
//
//	return sum
//}
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
//
//func timeTrack(start time.Time, name string) {
//	elapsed := time.Since(start)
//	log.Printf("%s took %s", name, elapsed)
//}
