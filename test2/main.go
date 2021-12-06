package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{1, 4, 1}, []int{1, 5, 1}))
	fmt.Println(Solution([]int{4, 4, 2, 4}, []int{5, 5, 2, 5}))
	fmt.Println(Solution([]int{2, 3, 4, 2}, []int{2, 5, 7, 2}))
	fmt.Println(Solution([]int{2, 3, 4}, []int{2, 5, 1, 2}))
}

func Solution(P []int, S []int) int {

	var totalPeople int
	for _, person := range P {
		totalPeople = totalPeople + person
	}

	sort.Sort(sort.Reverse(sort.IntSlice(S)))

	cars := 0
	for _, seat := range S {
		totalPeople = totalPeople - seat
		cars++
		if totalPeople <= 0 {
			break
		}
	}

	return cars
}