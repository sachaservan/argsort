package argsort

import (
	"testing"

	"github.com/ncw/gmp"
)

func TestIntArgSort(t *testing.T) {
	test := []*gmp.Int{
		gmp.NewInt(1),
		gmp.NewInt(25),
		gmp.NewInt(3),
		gmp.NewInt(5),
		gmp.NewInt(4),
	}

	sort := NewIntArgsort(test)
	sorted := SortIntsByArray(test, sort)
	t.Logf("sorted: %v", sorted)

	isSorted := true
	for i := range sorted {
		if i+1 >= len(sorted) {
			break
		}

		if sorted[i].Cmp(sorted[i+1]) > 0 {
			isSorted = false
			break
		}
	}

	if !isSorted {
		t.Fatalf("array not sorted!")
	}
}
