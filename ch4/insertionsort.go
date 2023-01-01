package main

import (
	"fmt"
	"math"
)

func InsertionSortInc(in []int) []int {
	out := make([]int, len(in)+1)
	out[0] = math.MaxInt
	for i, x := range in {
		var j int
		for j = 0; x >= out[j]; j++ {}
		for k := i; k >= j; k-- {
			out[k+1] = out[k]
		}
		out[j] = x
	}
	return out[:len(in)]
}

func InsertionSortDec(in []int) []int {
	out := make([]int, len(in)+1)
	out[0] = math.MinInt
	for i, x := range in {
		var j int
		for j = i+1; out[j-1] > x; j-- {
			out[j] = out[j-1]
		}
		out[j] = x
	}
	return out[1:]
}

func main()  {
	fmt.Println(InsertionSortInc([]int{2,432,3,43,5,56,2,313,5,4,6}))
	fmt.Println(InsertionSortDec([]int{2,432,3,43,5,56,2,313,5,4,6}))
}