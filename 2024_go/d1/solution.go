// --- Day 1: Historian Hysteria ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(file *os.File) (arr1 []int, arr2 []int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "   ")
		num, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			panic(err)
		}
		arr1 = append(arr1, num)

		num, err = strconv.Atoi(string(parts[1]))
		if err != nil {
			panic(err)
		}
		arr2 = append(arr2, num)
	}

	return arr1, arr2
}

func part1(left []int, right []int) {
	slices.Sort(left)
	slices.Sort(right)

	// distance computation
	// since the arrays are sorted
	// we only need to compare both numbers in the same index
	total := 0
	for i := 0; i < len(left) && i < len(right); i++ {
		num1 := left[i]
		num2 := right[i]

		if num1 > num2 {
			total += (num1 - num2)
			continue
		}

		total += (num2 - num1)
	}

	fmt.Println("Part1:", total)
}

func part2(left []int, right []int) {
	total := 0

	// sort the 2nd array
	// check how many times the num in the arr1
	// appears in arr2
	// since arr2 is sorted we can stop the inner loop
	// if num2 is greater than num1
	slices.Sort(right)
	for _, num1 := range left {
		similarity := 0

	inner:
		for _, num2 := range right {
			if num1 == num2 {
				similarity++
			}

			if num2 > num1 {
				break inner
			}

		}

		total += (num1 * similarity)
	}

	fmt.Println("Part2:", total)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	arr1, arr2 := parseInput(file)
	part1(arr1, arr2)
	part2(arr1, arr2)
}
