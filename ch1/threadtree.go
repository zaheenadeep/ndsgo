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
	q.RSon, q.RThread = p.RSon, p.RThread
	p.RSon, p.RThread = q, false // q becomes p's right son
	q.LSon, q.LThread = p, true // p becomes q's in-order pred
	if !q.RThread { // p originally had a right son 
		InOrderSuccessor(q).LSon = q // set q as in-order pred
	}
}