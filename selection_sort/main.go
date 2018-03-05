package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a []int
	size := 30
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(size))
	}
	fmt.Printf("%v\n", a)
	selsort(a)
	fmt.Printf("%v\n", a)
}

func selsort(a []int) {
	for i := range a {
		minIdx := i
		for j := i; j < len(a); j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}
