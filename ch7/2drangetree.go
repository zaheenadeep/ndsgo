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
	Mid		int
	Point	Point
	RTree	*RNode
	Left	*RNode
	Right	*RNode
}

func (t *RNode) RangeSearch2D(bX, eX, bY, eY int) []Point {

}

func (t *RNode) rangeSearch1D(b, e int) []Point {
	// tree traversal
	for t != nil && t.Type == NonLeaf {
		if b <= t.Mid {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	if t == nil { // t beyond last leaf node
		return nil
	}

	// list traversal
	if t.Mid < b {
		t = t.Right
	}
	var s []Point
	for t != nil && t.Mid <= e {
		s = append(s, t.Point)
		t = t.Right
	}
	return s
}