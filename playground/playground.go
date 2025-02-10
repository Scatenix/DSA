package main

import (
	"dsa/datastructures/linkedList"
	"fmt"
)

func main() {
	ll := linkedList.NewLinkedList[int]([]int{2, 6, 426, 56}...)

	ll.Remove(3)
	ll.Remove(1)
	ll.Remove(0)
	ll.Remove(0)

	fmt.Println(*ll)
}
