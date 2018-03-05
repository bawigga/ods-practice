package main

import (
	"math/rand"
)

func main() {
	var a []int
	size := 10
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(size))
	}
	mergesort(a)
}

func mergesort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a)
	left := mergesort(a[:mid])
	right := mergesort(a[mid+1:])
	size := len(left) + len(right)
	aux := make([]int, size)
	i, j, k := 0, 0, 0
	for k < size {
		if i < len(left)-1 && left[i] > right[j] {
			aux[k++] = left[i++]
		}
	}
}
