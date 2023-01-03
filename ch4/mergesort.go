package main

import "fmt"

func TopDownMergeSort(s []int) []int {
	if len(s) == 1 {
		return []int{s[0]}
	}
	
	m := (len(s)-1) / 2
	ls := TopDownMergeSort(s[:m+1])
	rs := TopDownMergeSort(s[m+1:])
	
	news := make([]int, len(s))
	i, j, k := 0, 0, 0
	for i < len(ls) && j < len(rs) {
		if ls[i] <= rs[j] {
			news[k] = ls[i]
			i++
		} else {
			news[k] = rs[j]
			j++
		}
		k++
	}
	if i < len(ls) {
		copy(news[k:], ls[i:])
	} else if j < len(rs) {
		copy(news[k:], rs[j:])
	}
	return news
}


func main()  {
	fmt.Println(TopDownMergeSort([]int{2,432,3,43,5,56,2,313,5,4,6}))
}