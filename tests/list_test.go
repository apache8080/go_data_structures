package list_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/apache8080/go_data_structures/list"
)

func TestAppend(t *testing.T) {
	l := list.NewList()

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
	l := list.NewList()

	for i := 0; i < 5; i++ {
		l.Append(i)
	}

	l.Remove(2)
	l.Remove(4)

	arr := l.GetArray()
	
	assert.Equal(t, len(arr), 3, "array size should be 3");

	assert.Equal(t, arr[0], 0, "arr[0] should equal 0");
	assert.Equal(t, arr[1], 2, "arr[0] should equal 0");
	assert.Equal(t, arr[2], 3, "arr[0] should equal 0");
}
