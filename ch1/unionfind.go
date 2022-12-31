package main

import "fmt"

type EqClass struct {
	Data	int
	Father	*EqClass
}

type EqRelation struct { A, B *EqClass }

func Union(p, q *EqClass) *EqClass {
	if p != nil && p.Data != q.Data {
		p.Father = q
	}
	return q
}

func Find(p *EqClass) *EqClass {
	// find the root r
	r := p
	for r.Father != nil {
		r = r.Father
	}

	// short-circuit the path by linking ec to r
	var father *EqClass
	for ec := p; ec.Father != nil; ec = father {
		father = ec.Father
		ec.Father = r
	}

	return r
}

func MakeSet(e int) *EqClass {
	return &EqClass{e, nil}
}

func UnionFind(rels []*EqRelation) {
	for _, rel := range rels {
		a, b := Find(rel.A), Find(rel.B)
		if a != b {
			Union(a, b)
		}
	}
}

func main() {
	var (
		One			= MakeSet(1)
		Two			= MakeSet(2)
		Three		= MakeSet(3)
		Four		= MakeSet(4)
		Five		= MakeSet(5)
		Seven		= MakeSet(7)
		Twelve		= MakeSet(12)
		ThirtyTwo	= MakeSet(32)
		ThirtyFour	= MakeSet(34)
		SixtyFour	= MakeSet(64)
		SeventyFive	= MakeSet(75)
		NinetyNine	= MakeSet(99)
	)
	rels := []*EqRelation{
		{One, Two}, {Three, Two}, {Four, Three}, // equivalent
		{SeventyFive, Twelve}, {ThirtyTwo, SixtyFour}, {ThirtyFour, NinetyNine}, // distinct
		{Seven, Seven}, {Seven, Five}, // equivalent
	}
	UnionFind(rels)

	// Tests
	fmt.Println(Find(Two) == Find(Three)) // true; Two and One equivalent
	fmt.Println(Find(SeventyFive) == Find(Twelve)) // true; SeventyFive and Twelve equivalent
	fmt.Println(Find(Two) == Find(Twelve))  // false; Two and Twelve not equivalent
	fmt.Println(Find(Seven) == Find(MakeSet(1000))) // false: Seven and new class not equivalent
}