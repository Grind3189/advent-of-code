// --- Day 2: Red-Nosed Reports ---
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
	var safe int

	for scanner.Scan() {
		const MIN, MAX = 1, 3

		text := scanner.Text()
		input := strings.Split(text, " ")

		result := checkArr(input)
		if result == false {
			continue
		}

		safe++
	}

	fmt.Println("Part 1:", safe)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)
	var unsafes [][]string
	var safe int

outer:
	for scanner.Scan() {
		text := scanner.Text()
		input := strings.Split(text, " ")

		result := checkArr(input)

		if result == true {
			safe++
			continue outer
		}

		unsafes = append(unsafes, input)

	}

unsafes:
	for _, input := range unsafes {
		for i := range input {
			newArr := removeEl(input, i)
			result := checkArr(newArr)
			if result == true {
				safe++
				continue unsafes
			}
		}
	}

	fmt.Println("Part 2:", safe)
}

func removeEl(arr []string, i int) []string {
	temp := make([]string, 0)
	temp = append(temp, arr[:i]...)
	return append(temp, arr[i+1:]...)
}

func checkArr(input []string) bool {
	const MIN, MAX = 1, 3
	asc, desc, prev := false, false, 0
check:
	for i, txt := range input {
		diff := 0
		num, err := strconv.Atoi(txt)
		if err != nil {
			panic(err)
		}

		if i == 0 {
			prev = num
			continue check
		}

		// each outer loop should only be asc or desc
		// to be considered valid
		if num > prev && desc == false {
			asc = true
			diff = num - prev
		} else if num < prev && asc == false {
			desc = true
			diff = prev - num
		} else {
			return false
		}

		prev = num

		// the difference between each level should be between
		// the minimum and maximum value to be considered valid
		if diff < MIN || diff > MAX {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	part1(file)

	file, err = os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	part2(file)
}
