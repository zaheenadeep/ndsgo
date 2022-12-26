package main

import "fmt"

type ThreadNode struct {
	Data		int
	LSon		*ThreadNode
	LThread		bool
	RSon		*ThreadNode
	RThread		bool
}

func InOrderSuccessor(p *ThreadNode) *ThreadNode {
	q := p.RSon
	if !p.RThread {
		// not a thread
		for !q.LThread {
			q = q.LSon
		}
	}
	return q
}

func AddNodeThreaded(p, q *ThreadNode) {
	q.RSon, q.RThread = p.RSon, p.RThread // p.RSon becomes q.RSon
	p.RSon, p.RThread = q, false // q becomes p.RSon
	q.LSon, q.LThread = p, true // p becomes in-order pred of q
	if !q.RThread { // p did have a right son originally
		InOrderSuccessor(q).LSon = q // set q's in-order pred
	}
}