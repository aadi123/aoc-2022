package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() {
	// read input
	f, err := os.Open("./input/day3.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPriorityScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		halfLine := line[:len(line)/2]
		chars := map[rune]struct{}{}
		for _, ch := range halfLine {
			chars[ch] = struct{}{}
		}

		// iterate over second half of string
		secondHalfLine := line[len(line)/2:]
		for _, ch := range secondHalfLine {
			if _, ok := chars[ch]; ok {
				if unicode.IsUpper(ch) {
					totalPriorityScore += int(ch-'A') + 27
				} else {
					totalPriorityScore += int(ch-'a') + 1
				}
				break
			}
		}
	}
}

func part2() {
	// read input
	f, err := os.Open("./input/day3.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPriorityScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		chars := map[rune]struct{}{}

		// map first line
		for _, ch := range line {
			chars[ch] = struct{}{}
		}

		tempMap := map[rune]struct{}{}
		if ok := scanner.Scan(); !ok {
			panic("Not in groups of 3")
		}
		for _, ch := range scanner.Text() {
			if _, ok := chars[ch]; ok {
				tempMap[ch] = struct{}{}
			}
		}
		chars = tempMap

		tempMap = map[rune]struct{}{}
		if ok := scanner.Scan(); !ok {
			panic("Not in groups of 3")
		}
		for _, ch := range scanner.Text() {
			if _, ok := chars[ch]; ok {
				tempMap[ch] = struct{}{}
			}
		}
		chars = tempMap
		for ch := range chars {
			if unicode.IsUpper(ch) {
				totalPriorityScore += int(ch-'A') + 27
			} else {
				totalPriorityScore += int(ch-'a') + 1
			}
		}
	}

	fmt.Println(totalPriorityScore)

}

func main() {
	part1()
	part2()
}
