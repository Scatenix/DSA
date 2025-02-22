package hashMap

import "dsa/algorithms/hash"

/*
Some ideas for my HashMap

Dynamic resizing:
	Resizing is typically done by doubling the size of the underlying array and recalculating every key.
	TODO: Think about adding a retrival of all nodes per linked list as function within the linked list itself or do it in this hashMap with node traversal.
*/

func main() {
	println("HashMap under construction!")
	println(hash.Murmur3("test", 2))
}
