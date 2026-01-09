// --- Day 5: Print Queue ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns page ordering and page numbers
func input(file *os.File) (pO [][]int, pN [][]int) {
	scanner := bufio.NewScanner(file)

	section := 1
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			section = 2
			continue
		}

		if section == 1 {
			nums := make([]int, 2)
			str := strings.Split(text, "|")
			num1, _ := strconv.Atoi(str[0])
			num2, _ := strconv.Atoi(str[1])
			nums[0] = num1
			nums[1] = num2
			pO = append(pO, nums)
			continue
		}

		// section 2
		strArr := strings.Split(text, ",")
		var nums []int
		for _, str := range strArr {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		pN = append(pN, nums)

	}

	return pO, pN
}

func part1(pO [][]int, pN [][]int) {
	for _, updates := range pN {
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pO, pN := input(file)
	fmt.Println("PO:", pO)
	fmt.Println("PN:", pN)
}
