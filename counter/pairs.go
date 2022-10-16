package counter

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].First < p[j].First
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
