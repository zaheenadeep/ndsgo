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
	s := NewSet()
	for _, rel := range rels {
		a := Find(rel.A)
		b := Find(rel.B)
		s.Remove(a)
		s.Remove(b)
		if a != b {	
			s.Add(Union(a, b))
		} else {
			s.Add(a)
		}
	}

	var keys []*EqClass
	for k := range s.list {
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

// Ignore everything after this.
// It is an abstraction defined for set operations using built-in go maps.
// from https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80

type Set struct {
	list map[*EqClass]struct{} //empty structs occupy 0 memory
}

func (s *Set) Has(v *EqClass) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v *EqClass) {
	s.list[v] = struct{}{}
}

func (s *Set) Remove(v *EqClass) {
	delete(s.list, v)
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[*EqClass]struct{})
	return s
}