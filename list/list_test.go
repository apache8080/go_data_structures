package list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	l := NewList()

	for i := 0; i < 5; i++ {
		l.Append(i)
	}

	arr := l.GetArray()

	assert.Equal(t, len(arr), 5, "array size should be 5")

	for i := 0; i < len(arr); i++ {
		assert.Equal(t, arr[i], i, "value in array should be correct at certain id")
	}
}

func TestRemove(t *testing.T) {
	l := NewList()

	for i := 0; i < 5; i++ {
		l.Append(i)
	}

	l.Remove(2)
	l.Remove(4)

	arr := l.GetArray()
	
	assert.Equal(t, len(arr), 3, "array size should be 3");

	assert.Equal(t, arr[0], 0, "arr[0] should equal 0");
	assert.Equal(t, arr[1], 2, "arr[1] should equal 2");
	assert.Equal(t, arr[2], 3, "arr[2] should equal 3");
}

func TestFindRemove(t *testing.T) {
	l := NewList()

	for i := 0; i < 5; i++ {
		l.Append(i)
	}

	l.FindRemove(2)
	l.FindRemove(4)

	arr := l.GetArray()
	
	assert.Equal(t, len(arr), 3, "array size should be 3");

	assert.Equal(t, arr[0], 0, "arr[0] should equal 0");
	assert.Equal(t, arr[1], 1, "arr[1] should equal 1");
	assert.Equal(t, arr[2], 3, "arr[2] should equal 3");
}

func TestFront(t *testing.T) {
	l := NewList()

	for i := 0; i < 5; i++ {
		l.Append(i)
	}

	assert.Equal(t, l.Front(), 0, "head pointer should equal 0");
}

func TestBack(t *testing.T) {
	l := NewList()

	for i := 0; i < 5; i++ {
		l.Append(i)
	}

	assert.Equal(t, l.Back(), 4, "tail pointer should equal 4");
}
