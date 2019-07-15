package score

type Range struct {
	Start int
	End   int
}

type Ranges []*Range

func (r Ranges) Len() int {
	return len(r)
}

func (r Ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Ranges) Less(i, j int) bool {
	return r[i].Start < r[j].Start
}
