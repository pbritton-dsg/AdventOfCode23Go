package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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

var moveDir = map[string]direction{
	"-1, 0": UP,
	"0, -1": LEFT,
	"0, 1":  RIGHT,
	"1, 0":  DOWN,
}

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

var nextDir = map[direction]direction{
	RIGHT: DOWN,
	LEFT:  UP,
	DOWN:  RIGHT,
	UP:    LEFT,
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		line := scanner.Text()
		for _, r := range line {
			n := int(r - '0')
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	fmt.Println(calcHeatLoss(grid, 1, 3))
	fmt.Println(calcHeatLoss(grid, 4, 10))
}

func calcHeatLoss(grid [][]int, min, max int) int {
	queue, seen := PQ[State]{}, map[State]struct{}{}
	queue.GPush(State{row: 0, col: 0, dir: RIGHT}, 0)
	queue.GPush(State{row: 0, col: 0, dir: DOWN}, 0)

	for len(queue) > 0 {
		state, heat := queue.GPop()

		if state.row == len(grid)-1 && state.col == len(grid[0])-1 {
			return heat
		}

		if _, ok := seen[state]; ok {
			continue
		}
		seen[state] = struct{}{}

		for i := -max; i <= max; i++ {
			nextMove, _ := curDir[state.dir]
			nr := state.row + (i * nextMove[0])
			nc := state.col + (i * nextMove[1])

			if (nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0])) || i > -min && i < min {
				continue
			}
			//Calculate heat along the path to grid[nr][nc]
			h, s := 0, int(math.Copysign(1, float64(i)))
			curRow, curCol := state.row, state.col
			for curRow != nr {
				curRow = curRow + (1 * s)
				h += grid[curRow][curCol]
			}
			for curCol != nc {
				curCol = curCol + (1 * s)
				h += grid[curRow][curCol]
			}

			next, _ := nextDir[state.dir]
			queue.GPush(State{nr, nc, next}, heat+h)
		}
	}
	return -1
}
