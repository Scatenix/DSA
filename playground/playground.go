package main

import (
	"dsa/datastructures/linkedList"
	"fmt"
)

func main() {
	ll := linkedList.NewLinkedList[int]([]int{2, 6, 426}...)

	fmt.Println(*ll)
}
