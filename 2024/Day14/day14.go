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

type Robot struct {
	x  int
	y  int
	vX int
	vY int
}

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	//grid := make([][]rune, 0)
	robots := make([]Robot, 0)
	maxX, maxY := 0, 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(strings.Split(fields[0][2:], ",")[0])
		y, _ := strconv.Atoi(strings.Split(fields[0][2:], ",")[1])
		vX, _ := strconv.Atoi(strings.Split(fields[1][2:], ",")[0])
		vY, _ := strconv.Atoi(strings.Split(fields[1][2:], ",")[1])
		robots = append(robots, Robot{x: x, y: y, vX: vX, vY: vY})

		maxX = max(x, maxX)
		maxY = max(y, maxY)
	}

	fmt.Println(partOne(robots, maxX, maxY, 100))
	//fmt.Println(partTwo(garden))
}

func partOne(robots []Robot, maxX, maxY, seconds int) int {
	defer timeTrack(time.Now(), "partOne")

	fmt.Println(robots[0].x, robots[0].y)
	for seconds > 0 {
		for i, _ := range robots {
			if robots[i].x+robots[i].vX >= 0 {
				robots[i].x = (robots[i].x + robots[i].vX) % (maxX + 1)
			} else {
				robots[i].x = (robots[i].x + robots[i].vX) + maxX + 1
			}
			if robots[i].y+robots[i].vY >= 0 {
				robots[i].y = (robots[i].y + robots[i].vY) % (maxY + 1)
			} else {
				robots[i].y = (robots[i].y + robots[i].vY) + maxY + 1
			}
		}
		seconds--
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.x < maxX/2 && robot.y < maxY/2 {
			q1++
		} else if robot.x > maxX/2 && robot.y < maxY/2 {
			q2++
		} else if robot.x < maxX/2 && robot.y > maxY/2 {
			q3++
		} else if robot.x > maxX/2 && robot.y > maxY/2 {
			q4++
		}
	}

	fmt.Println(q1, q2, q3, q4)

	return q1 * q2 * q3 * q4
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
