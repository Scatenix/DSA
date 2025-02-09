package main

import (
	"dsa/datastructures/linkedList"
	"fmt"
)

func main() {
	ll := linkedList.NewLinkedList[int]([]int{2, 6, 426}...)

	ll.Insert(5, 3)
	ll.Insert(1, 0)
	ll.Insert(10, 2)

	fmt.Println(*ll)
}
