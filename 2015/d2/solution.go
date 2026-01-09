// --- Day 2: I Was Told There Would Be No Math ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	totalSf := 0 // total square feet

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "x")
		var dimensions [3]int

		// conversion
		for i, char := range parts {
			num, _ := strconv.Atoi(char)
			dimensions[i] = num
		}

		// length, width, height
		l, w, h := dimensions[0], dimensions[1], dimensions[2]
		// sides of the box
		sides := [3]int{l * w, w * h, h * l}
		smallest := sides[0]
		for _, num := range sides {
			if num < smallest {
				smallest = num
			}
		}

		// calculate the total square feet
		sF := (2 * sides[0]) + (2 * sides[1]) + (2 * sides[2]) + smallest
		totalSf += sF
	}

	fmt.Println("Part 1:", totalSf)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)
	var total int

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "x")
		var dimensions [3]int
		// conversion
		for i, char := range parts {
			num, _ := strconv.Atoi(char)
			dimensions[i] = num
		}

		// find the largest side of the box
		largest, largestPos := dimensions[0], 0
		for i, num := range dimensions {
			if num > largest {
				largestPos = i
				largest = num
			}
		}

		// length, width, height
		l, w, h := dimensions[0], dimensions[1], dimensions[2]
		// feet of ribbon needed to wrap the present and
		// feet of ribbon needed to wrap the bow
		rFP, rFB := 0, (l * w * h)
	calculateRFP:
		for i := range dimensions {
			if largestPos == i {
				continue calculateRFP
			}
			rFP += (dimensions[i] + dimensions[i])
		}

		total += (rFB + rFP)
	}

	fmt.Println("Part 2:", total)
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
