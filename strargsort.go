package argsort

// Based on implementation of github.com/ericjster/
// https://gist.github.com/ericjster/1f44fda536728cbbfddd3df0e2a613d8

import (
	"sort"
)

// StringArgsort like in Numpy, it returns an array of indexes into an array. Note
// that the gonum version of argsort reorders the original array and returns
// indexes to reconstruct the original order.
type StringArgsort struct {
	s    []string // Points to orignal array but does NOT alter it.
	inds []int    // Indexes to be returned.
}

func (a StringArgsort) Len() int {
	return len(a.s)
}

func (a StringArgsort) Less(i, j int) bool {
	return a.s[a.inds[i]] < a.s[a.inds[j]]
}

func (a StringArgsort) Swap(i, j int) {
	a.inds[i], a.inds[j] = a.inds[j], a.inds[i]
}

// NewStringArgsort allocates and returns an array of indexes into the source float
// array.
func NewStringArgsort(src []string) []int {
	inds := make([]int, len(src))
	for i := range src {
		inds[i] = i
	}

	if len(src) != len(inds) {
		panic("floats: length of inds does not match length of slice")
	}
	a := &StringArgsort{s: src, inds: inds}
	sort.Sort(a)

	return inds
}

// SortStringsByArray sorts src based on inds
func SortStringsByArray(src []string, sort []int) []string {

	sorted := make([]string, 0)
	for i := range sort {
		sorted = append(sorted, src[sort[i]])
	}

	return sorted
}
