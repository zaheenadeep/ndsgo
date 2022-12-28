package main

import "fmt"

const MaxSons = 5 // arbitrary

type TNode struct {
	Data	int
	NumSons	int
	Sons	[MaxSons]*TNode
}

func main() {
	t := &TNode{
		Data:		1,
		NumSons:	3,
		Sons:		[MaxSons]*TNode{
			{Data: 5, NumSons: 0},
			{Data: 7, NumSons: 0},
			{Data: 9, NumSons: 0},
		},
	}
	fmt.Println(t)
}