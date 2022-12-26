package main

import "fmt"

type BNode struct {
	Data	int
	LSon	*BNode
	RSon	*BNode
}

func Similar1(t1, t2 *BNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else {
		return Similar1(t1.LSon, t2.LSon) &&
			   Similar1(t1.RSon, t2.RSon)
	}
}

func Similar2(t1, t2 *BNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else {
		return Similar2(t1.LSon, t2.LSon) &&
			   Similar2(t1.RSon, t2.RSon)
	}
}

func Similar3(t1, t2 *BNode) bool {
	if t1 == nil {
		return t2 == nil
	} else if t2 == nil {
		return false
	} else {
		return Similar3(t1.LSon, t2.LSon) &&
			   Similar3(t1.RSon, t2.RSon)
	}
}

func Equivalent(t1, t2 *BNode) bool {
	if t1 == nil {
		return t2 == nil
	} else if t2 == nil {
		return false
	} else {
		return t1.Data == t2.Data &&
			   Equivalent(t1.LSon, t2.LSon) &&
			   Equivalent(t1.RSon, t2.RSon)
	}
}

func InOrderStack(t *BNode) {
	var s []*BNode // stack declared as a slice
	b := t
	for !(b == nil && len(s) == 0) {
		if b != nil {
			s = append(s, b) // stack push
			b = b.LSon
		} else { // b == nil && len(s) > 0
			top := s[len(s)-1] // stack top
			visit(top.Data)
			s = s[:len(s)-1] // stack pop
			b = top.RSon
		}
	}
}

func visit(data int) {
	fmt.Println(data) // example for visit
} 

func InOrder(t *BNode) {
	if t == nil {
		return
	} else {
		InOrder(t.LSon)
		visit(t.Data)
		InOrder(t.RSon)
	}
}

func main() {
	fmt.Println(Similar3(nil, nil))
	fmt.Println(Similar3(&BNode{34, nil, nil}, nil))
	fmt.Println(Similar3(&BNode{34, nil, nil}, &BNode{34, nil, nil}))
	fmt.Println(Similar3(&BNode{34, nil, nil}, &BNode{56, nil, nil}))

	fmt.Println(Equivalent(nil, nil))
	fmt.Println(Equivalent(&BNode{34, nil, nil}, nil))
	fmt.Println(Equivalent(&BNode{34, nil, nil}, &BNode{34, nil, nil}))
	fmt.Println(Equivalent(&BNode{34, nil, nil}, &BNode{56, nil, nil}))
	fmt.Println(Equivalent(&BNode{34, nil, &BNode{56, nil, nil}}, &BNode{34, nil, &BNode{56, nil, nil}}))
	
	InOrderStack(&BNode{34, &BNode{23, nil, &BNode{27, nil, nil}}, &BNode{56, nil, nil}})
	InOrder(&BNode{34, &BNode{23, nil, &BNode{27, nil, nil}}, &BNode{56, nil, nil}})
}