type List struct {
	Data	int
	Next	*List
}

type Item struct {
	PredCount	int
	Successors	*List
}

