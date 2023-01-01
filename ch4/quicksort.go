package main

import (
	"fmt"
	"math/rand"
)

func Quicksort(s []int) {
	x := s[rand.Intn(len(s))]
	i, j := 0, len(s)-1
	for i <= j {
		for s[i] < x { i++ }
		for s[j] > x { j-- }
		if i <= j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	if 0 < j {
		Quicksort(s[:j+1])
	}
	if i < len(s) {
		Quicksort(s[i:])
	}
}

func main()  {
	s := []int{2,432,3,43,5,56,2,313,5,4,6}
	Quicksort(s)
	fmt.Println(s)
}