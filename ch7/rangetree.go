package rangetree

type RType int

const (
	Leaf RType = iota
	NonLeaf
)

type RNode {
	Type	RType
	Mid		int
	Left	*RNode
	Right	*RNode
}

func (t *Rnode) RangeSearch(b, e int) []int {
	// tree traversal
	for t != nil && t.Type == NonLeaf {
		if b <= t.Mid {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	if t == nil {
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