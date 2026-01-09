// --- Day 1: Not Quite Lisp ---
package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	cF := 0 // current floor

	for scanner.Scan() {
		text := scanner.Text()

		for _, runes := range text {
			char := string(runes)
			if char == "(" {
				cF++
			} else {
				cF--
			}

		}
	}

	fmt.Println("Part 1:", cF)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)
	cF := 0
	pos, inBasement := 0, false

	for scanner.Scan() {
		text := scanner.Text()

		for i, runes := range text {
			char := string(runes)
			if char == "(" {
				cF++
			} else {
				cF--
			}

			if cF == -1 && inBasement == false {
				inBasement = true
				pos = i + 1
			}
		}
	}

	fmt.Println("Part 2:", pos)
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
