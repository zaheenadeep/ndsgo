package main

import "fmt"

type RType int

const (
	Leaf RType = iota
	NonLeaf
)

type Point struct { X, Y int }

type RNode struct {
	Type	RType
	Mid		Point
	RTree	*RNode
	Left	*RNode
	Right	*RNode
}

func (t *RNode) RangeSearch2D(bx, ex, by, ey int) {
	
}

func (t *RNode) rangeSearch1D(b, e int) []int {
	// tree traversal
	for t != nil && t.Type == NonLeaf {
		if b <= t.Mid.X {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	if t == nil { // t beyond last leaf node
		return nil
	}

	// list traversal
	if t.Mid.X < b {
		t = t.Right
	}
	var s []int
	for t != nil && t.Mid.X <= e {
		s = append(s, t.Mid.X)
		t = t.Right
	}
	return s
}