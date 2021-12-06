package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{5}))
	fmt.Println(Solution([]int{23, 2, 4}))
	fmt.Println(Solution([]int{2, 3, 4, 2, 15, 9, 10}))
	fmt.Println(Solution([]int{2, 5, 23}))
	fmt.Println(Solution([]int{5, 19, 8, 1}))
}

type Factories []float64

func (f Factories) Len() int           { return len(f) }
func (f Factories) Less(i, j int) bool { return f[i] > f[j] }
func (f Factories) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func (fs Factories) Sum() float64 {
	var sum float64
	for _, f := range fs {
		sum = sum + f
	}
	return sum
}

func Solution(A []int) int {

	if len(A) == 1 {
		return 1
	}

	var (
		initialPollution float64
		currentPollution = make(Factories, len(A))
	)

	for i, pollution := range A {
		p := float64(pollution)
		currentPollution[i] = p
		initialPollution = initialPollution + p
	}

	goal := initialPollution / 2

	filters := 0
	sort.Sort(currentPollution)
	for {
		newFume := currentPollution[0] / 2
		currentPollution[0] = newFume

		filters++

		if currentPollution.Sum() <= goal {
			break
		}

		if currentPollution[0] < currentPollution[1] {
			sort.Sort(currentPollution)
		}

	}

	return filters
}
