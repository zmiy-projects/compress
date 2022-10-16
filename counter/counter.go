package counter

import (
	"sort"
	"strconv"
	"strings"
)

type Counter map[rune]int

func (c Counter) SortedSlice() []Pair {
	pairs := make(Pairs, 0, len(c))

	for key, value := range c {
		pairs = append(pairs, Pair{key, value})
	}

	sort.Sort(pairs)
	return pairs
}

func (c Counter) Add(value rune) {
	c[value]++
}

func (c Counter) Compress() string {
	builder := new(strings.Builder)

	for _, value := range c.SortedSlice() {
		builder.WriteRune(value.First)

		if value.Second > 1 {
			count := strconv.FormatInt(int64(value.Second), 10)
			builder.WriteString(count)
		}
	}

	return builder.String()
}

func CreateCounter(value string) *Counter {
	counter := Counter{}

	for _, word := range value {
		counter.Add(word)
	}

	return &counter
}
