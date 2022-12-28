package main

import "fmt"

type List struct {
	Data	int
	Next	*List
}

type Item struct {
	PredCount	int
	Successors	*List
}

type Relation struct{ I, J int }

func buildt(nt int, rels []Relation) []Item {
	t := make([]Item, nt)
	for _, r := range rels {
		t[r.J].PredCount++
		t[r.I].Successors = &List{r.J, t[r.I].Successors}
	}
	return t
}

func TopologicalSort(nt int, rels []Relation) []int {
	var s []int
	t := buildt(nt, rels)
	for i, item := range t {
		if item.PredCount == 0 {
			s = append(s, i)
		}
	}
	for i := 0; i < len(s); i++ {
		j := s[i]
		for l := t[j].Successors; l != nil; l = l.Next {
			succ := t[l.Data]
			succ.PredCount--
			if succ.PredCount == 0 {
				s = append(s, l.Data)
			}
		}
	}
	if len(s) != nt {
		return nil
	}
	return s
}

func main() {
	fmt.Println(TopologicalSort(7, []Relation{{3, 2}, {1, 0}, {0, 3}}))
}
