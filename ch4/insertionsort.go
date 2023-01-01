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

func main()  {
	fmt.Println(InsertionSortInc([]int{2,3,43,5,56,2,313,432,5,4,6}))
}