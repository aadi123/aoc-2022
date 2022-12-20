package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// declare individual elf calories and max Calories
	elfCals := 0
	maxCals := 0
	cals := 0

	// read input
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// while line is not new line, scan line and add to elfCals
	for scanner.Scan() {
		for scanner.Text() != "" {
			cals, _ = strconv.Atoi(scanner.Text())
			elfCals += cals
			scanned := scanner.Scan()
			if !scanned && scanner.Err() == nil {
				break
			}
		}

		if elfCals > maxCals {
			maxCals = elfCals
		}

		elfCals = 0
	}

	fmt.Println(maxCals)
}
