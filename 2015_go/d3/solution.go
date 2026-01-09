// --- Day 3: Perfectly Spherical Houses in a Vacuum ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	x, y, visited := 0, 0, 1
	loc := "0-0"
	locMap := map[string]int{loc: 1}

	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(text)

		for _, runes := range text {
			char := string(runes)

			// update x,y coordinates of santa
			switch char {
			case "^":
				x++
			case "v":
				x--
			case ">":
				y++
			case "<":
				y--
			}

			// check if the current loc from coordinates
			// has not been visited yet, update if not
			newLoc := strconv.Itoa(x) + "-" + strconv.Itoa(y)
			locVal := locMap[newLoc]
			if locVal == 0 {
				visited++
			}
			locMap[newLoc] = locVal + 1
			loc = newLoc
		}

	}
	fmt.Println("Part 1:", visited)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)
	// robo and santa x and y coordinates
	rX, rY, sX, sY := 0, 0, 0, 0
	visited := 1
	roboLoc, santaLoc := "0-0", "0-0"
	locMap := map[string]int{"0-0": 1}

	for scanner.Scan() {
		text := scanner.Text()

		for i, runes := range text {
			char := string(runes)
			var x, y *int
			var loc *string

			// determine who's moving
			// santa if even else its robo
			if i%2 == 0 {
				x, y = &sX, &sY
				loc = &santaLoc
			} else {
				x, y = &rX, &rY
				loc = &roboLoc
			}

			// update x,y coordinates
			switch char {
			case "^":
				*x++
			case "v":
				*x--
			case ">":
				*y++
			case "<":
				*y--
			}

			// check if the current loc from coordinates
			// has not been visited yet, update if not
			newLoc := strconv.Itoa(*x) + "-" + strconv.Itoa(*y)
			locVal := locMap[newLoc]
			if locVal == 0 {
				visited++
			}
			locMap[newLoc] = locVal + 1
			*loc = newLoc
		}
	}

	fmt.Println("Part 2:", visited)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	part1(file)

	file, err = os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	part2(file)
}
