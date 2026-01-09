// --- Day 4: Scratchcards ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type card struct {
	winningNums []string
	input       []string
}

func part1(cards []card) {
	var total int

	for _, card := range cards {

		winningNums := card.winningNums
		input := card.input

		var winningCount int

	counting:
		for _, wNum := range winningNums {
			// some input are blanks which is invalid
			if wNum == "" {
				continue counting
			}

			for _, num := range input {
				if num == wNum {
					// we dont need the winning numbers
					// what we need is how many times we will loop
					winningCount += 1
				}
			}
		}

		// e.g. winningCount = 4
		// 1, (1 * 2 = 2), (2 * 2 = 4), (4 * 2 = 8)
		// answer is: 8
		var points int
		const MULTIPLICATOR = 2
	calculations:
		for i := 0; i < winningCount; i++ {
			if i == 0 {
				points = 1
				continue calculations
			}
			points *= MULTIPLICATOR

		}
		total += points
	}

	fmt.Println("Part 1:", total)
}

func part2(cards []card) {
	// storing the index and no. of copies
	record := make(map[int]int)
	var total int

	for i, card := range cards {
		const ORIGINAL = 1

		winningNums := card.winningNums
		input := card.input
		matchingNums := 0
		// always add 1 original copy to the value each index
		record[i] += ORIGINAL

	counting:
		for _, wNum := range winningNums {
			// some input are blanks which is invalid
			if wNum == "" {
				continue counting
			}

			for _, num := range input {
				if num == wNum {
					// we dont need the winning numbers
					matchingNums += 1
				}
			}
		}

		// if we have 2 matching numbers
		// that means we will add 1 copy to the
		// 2 cards below this current card
		// e.g. current = index 1, target = index 2 and 3

		copies := record[i]
	tracking:
		for m := 1; m <= matchingNums; m++ {
			target := i + m

			// out of bounds
			if target > len(cards)-1 {
				break tracking
			}
			record[target] += copies
		}
	}

	// for the calculation
	// sum all the cards we have, these are both copies and original
	for _, val := range record {
		total += val
	}

	fmt.Println("Part 2:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards []card
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ":")
		list := strings.Split(parts[1], "|")
		// split the string into array of string
		// both for winning numbers and the numbers we have
		winningNums := strings.Split(strings.TrimSpace(list[0]), " ")
		input := strings.Split(strings.TrimSpace(list[1]), " ")
		cards = append(cards, card{winningNums, input})
	}

	part1(cards)
	part2(cards)
}
