package main

type List struct {
	Data	int
	Next	*List
}

type Item struct {
	PredCount	int
	Successors	*List
}

type Relation struct { I, J	int }

func buildt(nt int, rels []Relation) []Item {
	t := make([]Item, nt)
	for _, r := range rels {
		t[r.J].PredCount++
		t[r.I].Successors = &List{r.J, t[r.I].Successors}
	}
	return t
}

func TopologicalSort(nt int, rels []Relation) {
	var s []int
	t := buildt(nt, rels)
	for i, item := range t {
		if item.PredCount == 0 {
			s = append(s, i)
		}
	}
	for _, i := range s {
		for l := t[i].Successors; l != nil; l = l.Next {
			succ := t[l.Data]
			succ.PredCount--
			if succ.PredCount == 0 {
				s = append(s, l.Data)
			}
		}
	}
}

func main() {
	
}