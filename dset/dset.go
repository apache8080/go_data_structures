package dset

import (
	"errors"
)

type DisjointSet struct {
	set []int
}

func NewDisjointSet(size int) *DisjointSet {
	dset := &DisjointSet{}

	for i := 0; i < size; i++ {
		dset.set = append(dset.set, -1)
	}
	
	return dset
}

func (this *DisjointSet) Find(elem int) int {
	if elem < 0 || elem >= len(this.set) {
		return -1;
	}
	
	for this.set[elem] >= 0 {
		elem = this.set[elem]
	}
	
	return elem
}

func (this *DisjointSet) Union(a int, b int) error {
	if a < 0 || a >= len(this.set) || b < 0 || b > len(this.set) {
		return errors.New("Invalid input values.")
	}
	
	aVal := this.Find(a)
	bVal := this.Find(b)

	size := this.set[aVal] + this.set[bVal]
	
	if aVal != bVal {
		if this.set[aVal] >= this.set[bVal] {
			this.set[aVal] = size
			this.set[bVal] = aVal
		} else {
			this.set[bVal] = size
			this.set[aVal] = bVal 
		}
	}

	return nil;
}
