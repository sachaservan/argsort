package argsort

import (
	"testing"
)

func TestStringsArgSort(t *testing.T) {
	test := []string{
		"b",
		"a",
		"d",
		"d",
		"e",
	}

	sort := NewStringArgsort(test)
	sorted := SortStringsByArray(test, sort)
	t.Logf("sorted: %v", sorted)

	isSorted := true
	for i := range sorted {
		if i+1 >= len(sorted) {
			break
		}

		if sorted[i] > sorted[i+1] {
			isSorted = false
			break
		}
	}

	if !isSorted {
		t.Fatalf("array not sorted!")
	}
}
