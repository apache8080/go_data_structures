package list

import (
	"errors"
)

type Node struct {
	val int
	next *Node
}

type List struct {
	head *Node
	size int
}

func NewList() *List {
	return &List{nil, 0}
}

func (this *List) Append(val int) {
	if this.head == nil {
		this.head = &Node{val, nil}
	} else {
		curr := this.head
		for curr.next != nil {
			curr = curr.next
		}
		node := &Node{val, nil}
		curr.next = node
	}

	this.size++
}

func (this *List) GetArray() []int {
	arr := make([]int, this.size)

	curr := this.head
	
	for i := 0; i < this.size; i++ {
		arr[i] = curr.val
		curr = curr.next
	}

	return arr
}

func (this *List) Remove(n int) error {
	if (n < 1 ||  n > this.size) {
		return errors.New("Invalid element number.")
	} else if (n == 1) {
		curr := this.head.next
		this.head = curr
		this.size--
	} else if (n == this.size) {
		curr := this.head;
		for curr.next.next != nil {
			curr = curr.next
		}
		curr.next = nil
		this.size--
	} else {
		id := n - 1
		curr := this.head
		for i := 0; i < id - 1; i++ {
			curr = curr.next;
		}
		curr.next = curr.next.next
		this.size--
	}

	return nil
}

func (this *List) Find(val int) int {
	curr := this.head
	id := 0
	for curr != nil {
		if (curr.val == val) {
			return id + 1
		}
		curr = curr.next
		id++
	}

	return -1
}


