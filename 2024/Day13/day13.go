package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var visited [][]uint8

var neighbors = [][]int{
	{-1, 0}, {0, -1}, {0, 1}, {1, 0},
}

type Button struct {
	x int
	y int
}

type Prize struct {
	x int
	y int
}

type Game struct {
	buttonA Button
	buttonB Button
	prize   Prize
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	games := make([]Game, 0)
	for scanner.Scan() {
		buttonA := getButton(scanner.Text())
		scanner.Scan()
		buttonB := getButton(scanner.Text())
		scanner.Scan()
		prize := getPrize(scanner.Text())
		scanner.Scan()
		scanner.Text()

		games = append(games, Game{buttonA, buttonB, prize})
	}

	fmt.Println(partOne(games))
	//fmt.Println(partTwo(garden))
}

func getButton(line string) Button {
	nums := strings.Split(strings.Split(line, ":")[1], ",")
	x, _ := strconv.Atoi(nums[0][2:])
	y, _ := strconv.Atoi(nums[1][2:])
	return Button{x, y}
}

func getPrize(line string) Prize {
	nums := strings.Split(strings.Split(line, ":")[1], ",")
	x, _ := strconv.Atoi(strings.TrimSpace(nums[0][3:]))
	y, _ := strconv.Atoi(strings.TrimSpace(nums[1][3:]))
	return Prize{x, y}
}

func partOne(games []Game) int {
	defer timeTrack(time.Now(), "partOne")
	total := 0
	for _, game := range games {
		game.prize.x += 10000000000000
		game.prize.y += 10000000000000
		D := game.buttonA.x*game.buttonB.y - game.buttonB.x*game.buttonA.y
		Dx := game.prize.x*game.buttonB.y - game.buttonB.x*game.prize.y
		Dy := game.buttonA.x*game.prize.y - game.prize.x*game.buttonA.y

		if D != 0 && Dx == (Dx/D)*D && Dy == (Dy/D)*D {
			total += (Dx/D)*3 + (Dy / D)
		}
	}

	return total
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
