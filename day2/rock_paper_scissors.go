package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read input
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalSCore := 0
	var outcomeScore int
	var shapeScore int

	for scanner.Scan() {
		decisions := strings.Split(scanner.Text(), " ")
		var opponentShape string
		var myShape string
		switch decisions[0] {
		case "A":
			opponentShape = "Rock"
		case "B":
			opponentShape = "Paper"
		case "C":
			opponentShape = "Scissors"
		}

		switch decisions[1] {
		case "X":
			myShape = "Rock"
		case "Y":
			myShape = "Paper"
		case "Z":
			myShape = "Scissors"
		}

		if myShape == opponentShape {
			// draw
			outcomeScore = 3
		} else if (myShape == "Rock" && opponentShape == "Paper") || (myShape == "Paper" && opponentShape == "Scissors") || (myShape == "Scissors" && opponentShape == "Rock") {
			// lose
			outcomeScore = 0
		} else {
			outcomeScore = 6
		}

		switch myShape {
		case "Rock":
			shapeScore = 1
		case "Paper":
			shapeScore = 2
		case "Scissors":
			shapeScore = 3
		}

		totalSCore += shapeScore + outcomeScore
	}
	fmt.Println(totalSCore)
}
