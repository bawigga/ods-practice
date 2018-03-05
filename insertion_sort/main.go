package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a []int
	size := 30
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(size*2))
	}
	isort(a)
	fmt.Printf("%v", a)
}

func isort(a []int) {
	for i := range a {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}
