// --- Day 4: Ceres Search ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// parsed the input to
// ............ // added (index 0)
// ............ // added (1)
// ............ // added (2)
// ...XMASSD... // original first line
// ............ // added (last index - 2)
// ............ // added (last index - 1)
// ............ // added (last index)
// now we dont have to worry about out of bounds
func parseInput(file *os.File) (inputs []string) {
	const WORD = "XMAS"
	var boardLen int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		// append place holder at the first and last of text
		dummy := strings.Repeat(".", len(WORD)-1)
		text = dummy + text + dummy

		// add 3 rows w/ placeholder before the original index 0
		if len(inputs) == 0 {
			boardLen = len(text)
			dummy = strings.Repeat(".", boardLen)
			for range 3 {
				inputs = append(inputs, dummy)
			}
		}

		inputs = append(inputs, text)
	}

	// append 3 rows w/ placeholder at the last index
	for range 3 {
		inputs = append(inputs, strings.Repeat(".", boardLen))
	}
	return inputs
}

func part1(inputs []string) {
	var total int
	for i, input := range inputs {
	search:
		for j, rune := range input {
			char := string(rune)
			if char != "X" {
				continue search
			}

			// horizontal search
			h1 := input[j : j+4]
			h2 := input[j-3 : j+1]
			// vertical
			v1 := string(input[j]) + string(inputs[i-1][j]) + string(inputs[i-2][j]) + string(inputs[i-3][j])
			v2 := string(input[j]) + string(inputs[i+1][j]) + string(inputs[i+2][j]) + string(inputs[i+3][j])
			// diagonal l and r
			d1 := string(input[j]) + string(inputs[i-1][j-1]) + string(inputs[i-2][j-2]) + string(inputs[i-3][j-3])
			d2 := string(input[j]) + string(inputs[i-1][j+1]) + string(inputs[i-2][j+2]) + string(inputs[i-3][j+3])
			// diagonal bl and br
			d3 := string(input[j]) + string(inputs[i+1][j-1]) + string(inputs[i+2][j-2]) + string(inputs[i+3][j-3])
			d4 := string(input[j]) + string(inputs[i+1][j+1]) + string(inputs[i+2][j+2]) + string(inputs[i+3][j+3])

			patterns := []string{h1, h2, v1, v2, d1, d2, d3, d4}
			for _, p := range patterns {
				if p == "XMAS" || p == "SAMX" {
					total++
				}
			}
		}
	}

	fmt.Println("Part 1:", total)
}

func part2(inputs []string) {
	var total int
	for i, row := range inputs {
	search:
		for j, runes := range row {
			if string(runes) != "A" {
				continue search
			}

			// search diagonally
			d1 := string(inputs[i-1][j-1]) + string(runes) + string(inputs[i+1][j+1])
			d2 := string(inputs[i-1][j+1]) + string(runes) + string(inputs[i+1][j-1])

			if (d1 == "MAS" || d1 == "SAM") && (d2 == "MAS" || d2 == "SAM") {
				total++
			}
		}
	}

	fmt.Println("Part 2:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	inputs := parseInput(file)

	part1(inputs)
	part2(inputs)
}
