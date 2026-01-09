// --- Day 3: Gear Ratios ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// input is a modified
// we've concatinated "." at the first and last
// of each string in the input
func part1(input []string) {
	var total int
	for i, text := range input {
		// we'll use numStr for appending numbers
		// e.g. "12" + "3" "123"
		// then we will convert this later to actual number if this is a part number
		// we'll use txtToCheck to store every character sorrounding the number
		// then we will check it at once
		var numStr, txtToCheck string
		var start int
	inner:
		for j, rune := range text {

			if unicode.IsDigit(rune) {
				// store the starting index of the 1st number found
				if numStr == "" {
					start = j - 1
				}

				numStr += string(rune)
				continue inner
			}

			// at this point the current character in the loop is not a number
			// check if we have a value of number strings before
			// performing checks
			if numStr != "" {
				var before, after, up, down string
				valid := false

				// get left and right value of numbers
				// e.g. ".123*" = ".*"
				before = string(text[start])
				after = string(text[j])

				// get text row before this current text
				// e.g ".*..$" to ".*..$" // get this
				//     ".123."            // current text
				if i > 0 {
					up = input[i-1][start : j+1]
				}

				// get the bottom text after this text
				// e.g. ".123." // current text
				//      "....." // get this
				if i < len(input)-1 {
					down = input[i+1][start : j+1]
				}

				// append it all before checking
				txtToCheck = before + after + up + down
			check:
				for _, char := range txtToCheck {
					if char != '.' {
						valid = true
						break check
					}
				}

				if valid {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						panic("Unable to convert string to number")
					}
					total += num
				}

				// reset every operation
				numStr, txtToCheck, before, up, down, after = "", "", "", "", "", ""
				start = 0

			}
		}
	}

	fmt.Println("Part 1:", total)
}

// input is a modified
// we've concatinated "." at the first and last
// of each string in the input
func part2(input []string) {
	var total int

outer:
	for i, text := range input {
		if strings.Contains(text, "*") == false {
			// no asterisk found in this current text
			continue outer
		}

		// sorroundings is the value of
		// before, current and after index i
		var sorroundings []string
		// check up
		if i > 0 {
			sorroundings = append(sorroundings, input[i-1])
		}
		// bottom
		if i < len(input)-1 {
			sorroundings = append(sorroundings, input[i+1])
		}
		// current row
		sorroundings = append(sorroundings, input[i])

	findAsteriskI:
		for j, char := range text {
			if char != '*' {
				continue findAsteriskI
			}

			// when a number is on this
			// range it is considered valid
			min, mid, max := j-1, j, j+1
			var validNums []int

			// an asterisk has been found
			// loop through each of its sorrounding
			// to find a valid number
			for _, sorrounding := range sorroundings {
				var numStr string
				var isValidNum bool

			findNum:
				for k, char := range sorrounding {
					// out of max range
					if k > max && numStr == "" {
						break findNum
					}

					if unicode.IsDigit(char) {
						// validating if the current num
						// is adjacent to the asterisk
						if k == min || k == mid || k == max {
							isValidNum = true
						}
						numStr += string(char)
						continue findNum
					}

					if isValidNum == false {
						// num invalid continue finding
						// until we go out of range
						numStr = ""
						continue findNum
					}

					// at this point num is valid
					// means adjacent to asterisk
					// but we need exactly two of this
					num, err := strconv.Atoi(numStr)
					if err != nil {
						panic("Unable to convert num to string")
					}
					validNums = append(validNums, num)
					numStr = ""
					isValidNum = false
				}
			}

			// check if valid Gear
			// we need exactly two only
			if len(validNums) != 2 {
				continue findAsteriskI
			}

			num1, num2 := validNums[0], validNums[1]
			total += num1 * num2
		}

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
	// create an array of string
	// modify the each text, add '.' at beginning and end
	// this wont affect the value of each since '.' is not a symbol according to instruction
	// now we'll always have a value at the left and right side of the index
	// e.g. "123..%123" to ".123..%123."
	var input []string
	for scanner.Scan() {
		text := "." + scanner.Text() + "."
		input = append(input, text)
	}

	part1(input)
	part2(input)
}
