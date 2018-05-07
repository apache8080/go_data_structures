package dset

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	dset := NewDisjointSet(5)

	assert.Equal(t, dset.Find(0), 0, "Find(0) should be 0.")
}

func TestUnion(t *testing.T) {
	dset := NewDisjointSet(5)

	dset.Union(2, 3)

	assert.Equal(t, dset.Find(2), 2, "Find(2) should be 2")
	assert.Equal(t, dset.Find(3), 2, "Find(3) should be 2")
}

func TestMultiUnion(t *testing.T) {
	dset := NewDisjointSet(5)

	dset.Union(2, 3)
	dset.Union(0, 4)

	assert.Equal(t, dset.Find(2), 2, "Find(2) should be 2")
	assert.Equal(t, dset.Find(3), 2, "Find(3) should be 2")
	assert.Equal(t, dset.Find(0), 0, "Find(0) should be 0")
	assert.Equal(t, dset.Find(4), 0, "Find(4) should be 0")
}
