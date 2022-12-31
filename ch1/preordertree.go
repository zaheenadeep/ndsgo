package main

type LType int

const (
	Leaf LType = iota
	NonLeaf
)

type List struct {
	Data	int
	Type	LType
	Next	*List
}

type BNode struct {
	Data	int
	LSon	*BNode
	RSon	*BNode
}

func Reconstruct(l *List) *BNode {
	if l == nil {
		return nil
	}
	return Rebuild(l)
}

func Rebuild(l *List) *BNode {
	switch l.Type {
	case Leaf:
		return &BNode{l.Data, nil, nil}
	case NonLeaf:
		return &BNode{
			Data: l.Data,
			LSon: Rebuild(l.Next),
			RSon: Rebuild(l.Next.Next),
		}
	}
}