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
}