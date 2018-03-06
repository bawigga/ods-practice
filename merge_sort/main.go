package main

import (
	"fmt"
	"math/rand"
)

func main() {
	size := 20
	data := randArray(size)
	// mergesort(data)
	fmt.Printf("%v unsorted\n", data)
	fmt.Printf("%v   sorted\n", mergesort(data))
}

func randArray(size int) []int {
	var a []int
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(size))
	}
	return a
}

func mergesort(a []int) []int {
	// fmt.Printf("presort: %v\n", a)
	if len(a) <= 1 {
		return a
	}
	m := len(a) / 2
	left := mergesort(a[:m])
	right := mergesort(a[m:])
	sorted := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		sorted = append(sorted, left...)
	} else {
		sorted = append(sorted, right...)
	}
	// fmt.Printf("merged: %v\n", sorted)
	return sorted
}
