// --- Day 5: If You Give A Seed A Fertilizer ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type mapType struct {
	title  string
	values [][]int
}

func solve(seeds []int, maps []mapType) (lowestLoc int) {
search:
	for _, seed := range seeds {
		target := seed

		for _, mapList := range maps {
			numArrs := mapList.values
		numRange:
			for _, numArr := range numArrs {
				desNum := numArr[0]
				srcNum := numArr[1]
				rangeLen := numArr[2]
				maxSrc := srcNum + (rangeLen - 1)
				maxDes := desNum + (rangeLen - 1)

				// out of range
				if target < srcNum || target > maxSrc {
					continue numRange
				}

				diff := maxSrc - target
				target = maxDes - diff
				break numRange
			}
		}

		if lowestLoc == 0 {
			lowestLoc = target
			continue search
		}

		if target < lowestLoc {
			lowestLoc = target
		}

	}
	return lowestLoc
}

func part1(file *os.File) (seeds []int, maps []mapType) {
	scanner := bufio.NewScanner(file)
	var title string
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		seedsStr := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, seed := range seedsStr {
			num, err := strconv.Atoi(seed)
			if err != nil {
				panic(err)
			}
			seeds = append(seeds, num)
		}
		break
	}
parsing:
	for scanner.Scan() {
		text := scanner.Text()

		// skip blanks
		if text == "" {
			title = ""
			continue parsing
		}

		// get the title
		if strings.Contains(text, "map") {
			title = strings.Split(text, " ")[0]
			newMap := mapType{title: title}
			maps = append(maps, newMap)
			continue parsing
		}

		var nums []int
		numStr := strings.Split(text, " ")
		for _, str := range numStr {
			num, err := strconv.Atoi(str)
			if err != nil {
				continue parsing
			}
			nums = append(nums, num)
		}
	insert:
		for i, mapList := range maps {
			if mapList.title == title {
				maps[i].values = append(mapList.values, nums)
				break insert
			}
		}
	}

	return seeds, maps
}

func part2(seeds []int) []int {
	var seeds2 []int
	for i := 0; i < len(seeds); i = i + 2 {
		seed := seeds[i]
		rangeLen := seeds[i+1]

		for j := 0; j < rangeLen; j++ {
			seeds2 = append(seeds2, seed+j)
		}
	}

	return seeds2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// get the seeds for part 1
	seeds, maps := part1(file)

	start := time.Now()
	part1 := solve(seeds, maps)
	fmt.Println("Part 1:", part1)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	// get the seeds for part 2
	// start = time.Now()
	// seeds = part2(seeds)
	// part2 := solve(seeds, maps)
	// fmt.Println("Part 2:", part2)
	// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
