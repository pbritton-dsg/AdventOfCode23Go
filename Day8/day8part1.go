package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner.Scan()
	directions := scanner.Text()
	graph := Graph{nodes: make(map[string]Node)}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		node, left, right := line[0:3], line[7:10], line[12:15]
		newNode := Node{val: node, left: left, right: right}
		if _, exists := graph.nodes[node]; !exists {
			graph.nodes[node] = newNode
		}
	}

	fmt.Println(partOne(graph, directions))
}

func partOne(graph Graph, directions string) int {
	curNode, moves, i := graph.nodes["AAA"], 0, 0
	for curNode.val != "ZZZ" {
		if directions[i] == 82 { //Right byte
			curNode = graph.nodes[curNode.right]
		} else if directions[i] == 76 { //Left byte
			curNode = graph.nodes[curNode.left]
		}
		moves += 1
		i = (i + 1) % len(directions)

		if curNode.val == "ZZZ" {
			break
		}
	}

	return moves
}
