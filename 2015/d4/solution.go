// --- Day 4: The Ideal Stocking Stuffer ---
package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func part1(input string) {
	var fiveZeroes, newInput, result string

	count := 0
	for fiveZeroes != "00000" {
		count++
		// concatenating numbers to the default input
		newInput = input + strconv.Itoa(count)
		result = fmt.Sprintf("%x\n", md5.Sum([]byte(newInput)))
		// get the first five characters
		fiveZeroes = result[:5]
	}

	fmt.Println("Part 1:", count)
}

func part2(input string) {
	var sixZeroes, newInput, result string

	count := 0
	for sixZeroes != "000000" {
		count++
		// concatenating numbers to the default input
		newInput = input + strconv.Itoa(count)
		result = fmt.Sprintf("%x\n", md5.Sum([]byte(newInput)))
		// get the first five characters
		sixZeroes = result[:6]
	}

	fmt.Println("Part 1:", count)
}

func main() {
	input := "iwrupvqb"

	part1(input)
	part2(input)
}
