package argsort

// Based on implementation of github.com/ericjster/
// https://gist.github.com/ericjster/1f44fda536728cbbfddd3df0e2a613d8

import (
	"sort"

	"github.com/ncw/gmp"
)

// IntArgsort like in Numpy, it returns an array of indexes into an array. Note
// that the gonum version of argsort reorders the original array and returns
// indexes to reconstruct the original order.
type IntArgsort struct {
	s    []*gmp.Int // Points to orignal array but does NOT alter it.
	inds []int      // Indexes to be returned.
}

func (a IntArgsort) Len() int {
	return len(a.s)
}

func (a IntArgsort) Less(i, j int) bool {
	return a.s[a.inds[i]].Cmp(a.s[a.inds[j]]) < 0
}

func (a IntArgsort) Swap(i, j int) {
	a.inds[i], a.inds[j] = a.inds[j], a.inds[i]
}

// NewIntArgsort allocates and returns an array of indexes into the source float
// array.
func NewIntArgsort(src []*gmp.Int) []int {
	inds := make([]int, len(src))
	for i := range src {
		inds[i] = i
	}

	if len(src) != len(inds) {
		panic("floats: length of inds does not match length of slice")
	}
	a := &IntArgsort{s: src, inds: inds}
	sort.Sort(a)

	return inds
}

// SortIntsByArray sorts src based on inds
func SortIntsByArray(src []*gmp.Int, sort []int) []*gmp.Int {

	sorted := make([]*gmp.Int, 0)
	for i := range sort {
		sorted = append(sorted, src[sort[i]])
	}

	return sorted
}

// ReverseInts reverses the slice and returns it
func ReverseInts(src []*gmp.Int) []*gmp.Int {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}

	return src
}
