package main

import "fmt"

type RType int

const (
	Leaf RType = iota
	NonLeaf
)

type RNode struct {
	Type	RType
	Mid		int
	Left	*RNode
	Right	*RNode
}

func (t *RNode) RangeSearch(b, e int) []int {
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
	var s []int
	for t != nil && t.Mid <= e {
		s = append(s, t.Mid)
		t = t.Right
	}
	return s
}

func main() {
	t := &RNode{Leaf, 32, nil, nil}
	fmt.Println(t.RangeSearch(32, 45))
	t = &RNode{Leaf, 20, nil, nil}
	fmt.Println(t.RangeSearch(32, 45))

	c1 := &RNode{Leaf, 30, nil, nil}
	c2 := &RNode{Leaf, 50, c1, nil}
	c1.Right = c2
	t = &RNode{NonLeaf, 40, c1, c2}
	fmt.Println(t.RangeSearch(20, 45))
	fmt.Println(t.RangeSearch(20, 95))
	fmt.Println(t.RangeSearch(40, 95))
	fmt.Println(t.RangeSearch(95, 100))
}