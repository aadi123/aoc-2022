package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// print the elf with the most calories
// func part1() {
// 	// declare individual elf calories and max Calories
// 	elfCals := 0
// 	maxCals := 0
// 	cals := 0
// 	// read input
// 	f, err := os.Open("./input.txt")
// 	check(err)
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)

// 	// while line is not new line, scan line and add to elfCals
// 	for scanner.Scan() {
// 		for scanner.Text() != "" {
// 			cals, _ = strconv.Atoi(scanner.Text())
// 			elfCals += cals
// 			scanned := scanner.Scan()
// 			if !scanned && scanner.Err() == nil {
// 				break
// 			}
// 		}

// 		if elfCals > maxCals {
// 			maxCals = elfCals
// 		}

// 		elfCals = 0
// 	}

// 	fmt.Println(maxCals)

// }

// An IntHeap is a min-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func part2() {
	// declare individual elf calories and max Calories
	elfCals := 0
	cals := 0

	// read input
	f, err := os.Open("./input/day1.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	h := &MaxHeap{}
	heap.Init(h)

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

		heap.Push(h, elfCals)

		elfCals = 0
	}

	topThreeTotalCals := 0
	for i := 0; i < 3; i++ {
		topThreeTotalCals += heap.Pop(h).(int)
	}

	fmt.Println(topThreeTotalCals)
}

func main() {
	// part1()
	part2()
}
