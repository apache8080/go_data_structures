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
	tail *Node
	size int
}

func NewList() *List {
	return &List{nil, nil, 0}
}

func (this *List) Append(val int) {
	if this.head == nil {
		this.head = &Node{val, nil}
	} else {
		node := &Node{val, nil}
		if this.tail == nil {
			this.tail = node
			this.head.next = this.tail
		} else {
			temp := this.tail
			this.tail = node
			temp.next = this.tail
		}
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
		for curr.next != this.tail {
			curr = curr.next
		}
		curr.next = nil
		this.tail = curr
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

func (this *List) FindRemove(val int) {
	id := this.Find(val);
	if (id != -1) {
		this.Remove(id);
	}
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

func (this *List) Front() int {
	return this.head.val
}

func (this *List) Back() int {
	return this.tail.val
}
