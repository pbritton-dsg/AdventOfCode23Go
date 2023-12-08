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
	fmt.Println(partTwo(graph, directions))
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
	}

	return moves
}

func partTwo(graph Graph, directions string) int {
	curNodes := make([]Node, 0)
	totalMoves := make([]int, 0)
	for k, v := range graph.nodes {
		if k[2:] == "A" {
			curNodes = append(curNodes, v)
		}
	}

	for i := 0; i < len(curNodes); i++ {
		moves, dirIdx := 0, 0
		for curNodes[i].val[2:] != "Z" {
			if directions[dirIdx] == 82 { //Right byte
				curNodes[i] = graph.nodes[curNodes[i].right]
			} else if directions[dirIdx] == 76 { //Left byte
				curNodes[i] = graph.nodes[curNodes[i].left]
			}
			moves += 1
			dirIdx = (dirIdx + 1) % len(directions)
		}
		totalMoves = append(totalMoves, moves)
	}

	return lcm(totalMoves, 0)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(moves []int, idx int) int {
	if idx == len(moves)-1 {
		return moves[idx]
	}
	a := moves[idx]
	b := lcm(moves, idx+1)
	return a * b / gcd(a, b)
}
