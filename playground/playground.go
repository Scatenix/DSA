package main

import (
	"dsa/datastructures/hashMap"
	"dsa/util/perf"
	"time"
)

func main() {
	myMap()
	goMap()
}

func goMap() {
	hm := make(map[int]int)

	defer perf.MeasurePerformance()
	defer perf.TimeTracker(time.Now(), "go map")
	elements := 999999
	for elements > 0 {
		hm[elements] = elements
		elements--
	}
}

func myMap() {
	hm := hashMap.NewHashMap[int, int](0)

	defer perf.MeasurePerformance()
	defer perf.TimeTracker(time.Now(), "my map")
	elements := 999999
	for elements > 0 {
		hm.Insert(elements, elements)
		elements--
	}

	nilCount := 0
	for _, v := range hm.Pairs {
		if v == nil {
			nilCount++
		}
	}

	println(nilCount)

	longestLL := 0
	for _, v := range hm.Pairs {
		if v != nil {
			node := v.Head
			l := 0
			for node != nil {
				l++
				node = node.Next
			}
			if l > longestLL {
				longestLL = l
			}
		}
	}
	println(longestLL)
}
