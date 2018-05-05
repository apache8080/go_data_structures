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
