package main

import (
	"fmt"

	"github.com/apache8080/go_data_structures/list"
)

func main() {
	l := list.NewList()

	for i := 0; i < 5; i++ {
		l.Append(i)
	}

	arr := l.GetArray()

	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", arr[i])
	}

	fmt.Printf("Removing the 2nd element\n")
	l.Remove(2)

	fmt.Printf("Removing the 4th element\n")
	l.Remove(4)

	arr = l.GetArray()

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\n", arr[i])
	}

	l.Append(5)
	fmt.Printf("Appending 5 to the list\n")

	arr = l.GetArray()
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\n", arr[i])
	}	
}
