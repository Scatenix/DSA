package hashMap

import "dsa/algorithms/hash"

/*
Some ideas for my HashMap

Dynamic resizing:
- When resizing, check size. if size n is 2^32 then switch to a hash function which outputs 64 Bit
- Resize perhaps if a linked list has a size of >= 5 ? If we assume to have a hash function with good distribution, this should catch if we are overflowing.
	-> resize to size n + n * 10% (or something like that) and perhaps if n < 100, resize to n * 2
*/

func main() {
	println("HashMap under construction!")
	println(hash.Sha256("test"))
}
