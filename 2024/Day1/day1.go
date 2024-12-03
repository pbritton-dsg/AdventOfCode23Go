package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Peek() int {
	if h.Len() > 0 {
		return h[0]
	}
	return -1
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	leftList := &IntHeap{}
	heap.Init(leftList)
	rightList := &IntHeap{}
	heap.Init(rightList)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(fields[0])
		heap.Push(leftList, num)
		num, _ = strconv.Atoi(fields[1])
		heap.Push(rightList, num)
	}

	//fmt.Println(partOne(leftList, rightList))
	fmt.Println(partTwo(leftList, rightList))
}

func partOne(leftList *IntHeap, rightList *IntHeap) int {
	sum := 0
	for leftList.Len() > 0 {
		sum += abs(heap.Pop(leftList).(int) - heap.Pop(rightList).(int))
	}
	return sum
}

func partTwo(leftList *IntHeap, rightList *IntHeap) int {
	sum, leftNum, rightNum, rightNumCount, leftNumCount := 0, -1, -1, 1, 1
	popRight, popLeft := true, true
	for leftList.Len() > 0 && rightList.Len() > 0 {
		if popLeft {
			leftNum = heap.Pop(leftList).(int)
			for leftList.Peek() == leftNum {
				leftNumCount++
				heap.Pop(leftList)
			}
		}
		if popRight {
			rightNum = heap.Pop(rightList).(int)
			for rightList.Peek() == rightNum {
				rightNumCount++
				heap.Pop(rightList)
			}
		}

		if leftNum > rightNum {
			rightNumCount = 1
			popLeft, popRight = false, true
			continue
		} else if leftNum < rightNum {
			popLeft, popRight = true, false
			leftNumCount = 1
		} else {
			sum += leftNum * leftNumCount * rightNumCount
			leftNumCount, rightNumCount = 1, 1
			popLeft, popRight = true, true
		}
	}

	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
