package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// get the first and last digit in a text from a file
// then form a two digit number and get the sum
// e.g "1abc2" to 12
func part1(file *os.File) {
	var total int
	var numToAdd []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var numStrs []string
		text := scanner.Text()

		// get all available numbers in a line of text
		// put it in an array
		for _, char := range text {
			// skip if the unicode is not a number
			if unicode.IsNumber(char) == false {
				continue
			}

			// convert the unicode into string and append it
			str := string(char)
			numStrs = append(numStrs, str)
		}

		// get first and last string of number in the array and concatinate it
		// e.g ["1","3"]
		// concat = "13"
		num1 := numStrs[0]
		num2 := numStrs[len(numStrs)-1]
		concat := num1 + num2

		// convert the string to number
		// append it to the list of numbers to add
		num, err := strconv.Atoi(concat)
		if err != nil {
			panic("Unable to convert string to number")
		}
		numToAdd = append(numToAdd, num)
	}

	// get the sum of all numbers
	for _, num := range numToAdd {
		total += num
	}

	fmt.Println("Part 1:", total)
}

// get the first and last number from text in a file
// both word form and digit is valid
// "zoneight234" to "14"
func part2(file *os.File) {
	scanner := bufio.NewScanner(file)
	var numbers []int
	var total int

	for scanner.Scan() {
		line := scanner.Text()
		var numStr1 string
		var numStr2 string

		// check for a first number either digit or word form
	lToR:
		for i, char := range line {
			text := line[:i+1]
			numStr1 = checkWordForm(text)
			if numStr1 != "" {
				break lToR
			}

			if unicode.IsNumber(char) {
				numStr1 = string(char)
				break lToR
			}
		}

		// same as the loop above
		// but we'll check the text from right to left
	rToL:
		for i := len(line); i > 0; i-- {
			text := line[i:]
			char := rune(line[i-1])

			// same process with left to right
			// repeating it wont hurt instead of creating another function
			numStr2 = checkWordForm(text)
			if numStr2 != "" {
				break rToL
			}

			if unicode.IsNumber(char) {
				numStr2 = string(char)
				break rToL
			}
		}

		// concat both number string
		// "1" + "2" = "12"
		// parsed it to number and append to numbers arr
		concat := numStr1 + numStr2
		parsed, err := strconv.Atoi(concat)
		if err != nil {
			panic("Unable to convert string to number")
		}

		numbers = append(numbers, parsed)
	}

	// get the total of each number in the array to get the answer
	for _, num := range numbers {
		total += num
	}
	fmt.Println("Part 2:", total)
}

func checkWordForm(text string) string {
	if strings.Contains(text, "one") {
		return "1"
	} else if strings.Contains(text, "two") {
		return "2"
	} else if strings.Contains(text, "three") {
		return "3"
	} else if strings.Contains(text, "four") {
		return "4"
	} else if strings.Contains(text, "five") {
		return "5"
	} else if strings.Contains(text, "six") {
		return "6"
	} else if strings.Contains(text, "seven") {
		return "7"
	} else if strings.Contains(text, "eight") {
		return "8"
	} else if strings.Contains(text, "nine") {
		return "9"
	} else {
		return ""
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer file.Close()

	part1(file)

	file, err = os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer file.Close()

	part2(file)
}
