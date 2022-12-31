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

func UnionFind(rels []*EqRelation) []*EqClass {
	m := make(map[*EqClass]bool) // use built-in map as a set
	for _, rel := range rels {
		a := Find(rel.A)
		b := Find(rel.B)
		delete(m, a) // set remove
		delete(m, b) // set remove
		if a != b {	
			m[Union(a, b)] = true // set add
		} else {
			m[a] = true // set add
		}
	}

	var keys []*EqClass
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	var (
		Two		= MakeSet(2)
		Three	= MakeSet(3)
		Seven	= MakeSet(7)
	)
	rels := []*EqRelation{
		{MakeSet(1), Two}, {Three, Two}, {MakeSet(4), Three}, // equivalent
		{MakeSet(75), MakeSet(12)}, {MakeSet(32), MakeSet(64)}, {MakeSet(34), MakeSet(99)}, // distinct
		{Seven, Seven}, {Seven, MakeSet(5)}, // equivalent
	}

	uf := UnionFind(rels)
	for _, ec := range uf {
		fmt.Println(ec.Data)
	}
}