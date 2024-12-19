package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Used a priority queue implementation I found online since go doesn't have one.
type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

type State struct {
	row int
	col int
	dir direction
}

type direction int

const (
	UP    direction = 4
	DOWN            = 3
	LEFT            = 2
	RIGHT           = 1
)

var curDir = map[direction][]int{
	RIGHT: {0, 1},
	LEFT:  {0, -1},
	DOWN:  {1, 0},
	UP:    {-1, 0},
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]rune, 0)
	for scanner.Scan() {
		row := make([]rune, 0)
		line := scanner.Text()
		for _, r := range line {
			row = append(row, r)
		}
		grid = append(grid, row)
	}

	startRow, startCol := 0, 0
	for r, _ := range grid {
		for c, _ := range grid[r] {
			if grid[r][c] == 'S' {
				startRow, startCol = r, c
			}
		}
	}

	fmt.Println(findPath(grid, startRow, startCol))
	//fmt.Println(calcHeatLoss(grid, 4, 10))
}

func findPath(grid [][]rune, startRow, startCol int) int {
	queue, seen := PQ[State]{}, map[State]struct{}{}

	for d, dir := range curDir {
		nr, nc := startRow+dir[0], startCol+dir[1]
		if grid[nr][nc] != '#' {
			queue.GPush(State{row: startRow, col: startCol, dir: d}, 0)
		}
	}

	for len(queue) > 0 {
		state, cost := queue.GPop()

		if grid[state.row][state.col] == 'E' {
			return cost
		}

		if _, ok := seen[state]; ok {
			continue
		}
		seen[state] = struct{}{}

		for dir, neighbor := range curDir {
			nr, nc := state.row+neighbor[0], state.col+neighbor[1]
			if nr > -1 && nr < len(grid) && nc > -1 && nc < len(grid[state.row]) && grid[nr][nc] != '#' && grid[nr][nc] != 'S' {
				c := 1
				if state.dir != dir {
					c += 1000
				}

				queue.GPush(State{nr, nc, dir}, cost+c)
			}
		}
	}
	return -1
}
