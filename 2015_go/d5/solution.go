// --- Day 5: Doesn't He Have Intern-Elves For This? ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var total int
scan:
	for scanner.Scan() {
		text := scanner.Text()

		// it must have at least 3 vowels
		re := regexp.MustCompile("[aeiou]")
		matched := re.FindAllString(text, -1)
		if len(matched) < 3 {
			continue scan
		}

		pairs := getPairs(text)
		// it must not contain a blacklisted string
		blacklisted := []string{"ab", "cd", "pq", "xy"}
		hasAppearedTwice := false
	outer:
		for _, pair := range pairs {
			for _, blacklist := range blacklisted {
				if pair == blacklist {
					continue scan
				}
			}

			if hasAppearedTwice == true {
				continue outer
			}
			// it must have a 1 letter that appeared
			// twice in a row e.g. abcdde (dd)
			char1, char2 := pair[0], pair[1]
			if char1 == char2 {
				hasAppearedTwice = true
			}
		}

		// at this point the text has passed the
		// vowel and blacklist requirements
		if hasAppearedTwice == true {
			total++
		}
	}

	fmt.Println("Part 1:", total)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)
	var total int

input:
	for scanner.Scan() {
		text := scanner.Text()
		// it must contain a pair of two letter that
		// appears atleast twice without overlap
		// e.g. aaxaa but not aaa
		{
			pairs := getPairs(text)
			passed := false

		outer:
			for i, pTC := range pairs {
				if i == len(pairs)-1 {
					break
				}

				for j, pair := range pairs {
					// since getPairs has overlapping pairs
					// we will skip it
					if j == i || j == i+1 || j+1 == i {
						continue
					}
					if pTC == pair {
						passed = true
						break outer
					}
				}
			}
			if passed == false {
				continue input
			}
		}
		// it must contains 1 letter which repeats
		// with exactly 1 letter between them
		// e.g. xyx or aaa
		{
			for i := range text {
				if i == len(text)-2 {
					break
				}
				chars := text[i : i+3]

				if chars[0] == chars[2] {
					total++
					continue input
				}

			}
		}

	}
	fmt.Println("Part 2:", total)
}

// divide text into pairs with overlap
// e.g. "aba" to "ab", "ba"
func getPairs(text string) []string {
	var pairs []string
	for i := 0; i < len(text)-1; i++ {
		pair := string(text[i]) + string(text[i+1])
		pairs = append(pairs, pair)
	}
	return pairs
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
