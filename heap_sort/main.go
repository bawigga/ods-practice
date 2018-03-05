package main

import (
	"log"
	"math/rand"
)

func main() {
	var a []int
	size := 10000
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(size))
	}

	heapsort(a)

	for i := 0; i < size-1; i++ {
		if a[i] > a[i+1] {
			log.Fatalf("out of order!")
		}
	}

	log.Println("in order!")
}

func heapify(a []int) {
	log.SetPrefix("heapify: ")
	for i := 1; i < len(a); i++ {
		inPlace := false
		parent := (i - 1) / 2
		c := i
		for !inPlace {
			// log.Printf("e: %d, nidx: %d, parent: %d", a[c], c, a[parent])
			if a[c] <= a[parent] || c == 0 {
				inPlace = true
				break
			}
			// log.Printf("swapping %d with %d", a[c], a[parent])
			a[c], a[parent] = a[parent], a[c]
			c = parent
			parent = (c - 1) / 2
		}
	}
}

func heapsort(a []int) {
	heapify(a)
	// log.SetPrefix("sorting: ")
	n := len(a) - 1
	for i := n; i > 0; i-- {
		// log.Printf("status %v", a)
		// log.Printf("swapping i: %d %d %d", i, a[0], a[i])
		a[0], a[i] = a[i], a[0]
		// log.Printf("swapped %v", a)
		bubbleDown(a[:i])
	}
}

func bubbleDown(a []int) {
	li := len(a) - 1
	i := 0
	for {
		lc := 2*i + 1
		rc := lc + 1

		var largerChild int

		if lc < li && a[i] < a[lc] {
			largerChild = lc
		}

		if rc < li && a[i] < a[rc] && a[lc] < a[rc] {
			largerChild = rc
		}

		if largerChild > 0 {
			a[i], a[largerChild] = a[largerChild], a[i]
			i = largerChild
			continue
		}
		return
	}
}
