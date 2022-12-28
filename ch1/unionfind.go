package main

import "fmt"

type EqClass struct {
	Data	int
	Father	*EqClass
}

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
	for ec := p; ec != nil; ec = father {
		father = ec.Father
		ec.Father = r
	}

	return r
} 