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

func UnionFind(eqrels []*EqRelation) []*EqClass {
	m := make(map[*EqClass]bool)
	for _, rel := range eqrels {
		a := Find(rel.A)
		b := Find(rel.B)
		if a != b {
			m[Union(a, b)] = true
		} else {
			m[a], m[b] = true, true
		}
	}
	var s []*EqClass
	for k, _ := range m {
		s = append(s, k)
	}
	return s
}