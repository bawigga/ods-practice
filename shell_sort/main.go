package main

import (
	"fmt"
	"math/rand"
)

// tons of interesting research on picking good gaps for shell sort.
// the original author: Shell, used a simple gaps sequence that just did n/2 for
// each gap iteration. This leads to poor performance at the worst case of n**2.
// below are a few implementations of some more performant gap sequences.

var gapSequences = []string{
	"shell",
	"hibbard",
	"cuira",
	"knuth",
	"tokuda",
	"sedgewick",
	"sedgewick2",
}

func main() {
	size := 20
	data := randArray(size)

	for _, g := range gapSequences {
		tmp := make([]int, size)
		copy(tmp, data)
		shellsort(tmp, g)
	}
}

func randArray(size int) []int {
	var a []int
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(size))
	}
	return a
}

func shellsort(a []int, seq string) {
	gaps := gapSequence(seq, len(a))
	comparisons := 0
	for _, gap := range gaps {
		for i, j := 0, gap; j < len(a); i, j = i+1, j+1 {
			comparisons++
			for k := j; k-gap >= 0 && a[k] < a[k-gap]; k, comparisons = k-gap, comparisons+1 {
				a[k], a[k-gap] = a[k-gap], a[k]
			}
		}
	}
	fmt.Printf("comparisons: % 10s: % 10d\n", seq, comparisons)
}

func gapSequence(seq string, size int) []int {
	size /= 2
	var gaps []int
	switch seq {
	case "shell":
		return shellGaps(size)
	case "hibbard":
		// Hibbard Gap sequence
		// https://oeis.org/A168604
		gaps = []int{1, 3, 7, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095}
	case "cuira":
		// Ciura Gap Sequence
		// Best Increments for the Average Case of Shellsort, 13th International Symposium
		// on Fundamentals of Computation Theory, Riga, Latvia, 22-24 August 2001; Lecture
		// Notes in Computer Science 2001; Vol. 2138, pp. 106-117.
		// https://oeis.org/A102549
		gaps = []int{1, 4, 10, 23, 57, 132, 301, 701, 1750}
	case "knuth":
		gaps = []int{1, 4, 13, 40, 121, 364, 1093}
	case "tokuda":
		gaps = []int{1, 4, 9, 20, 46, 103, 233, 525, 1182, 2660, 5985, 13467}
	case "sloane":
		gaps = []int{1, 4, 13, 40, 121, 364, 1093, 3280, 9841, 29524}
	case "sedgewick":
		gaps = []int{1, 8, 23, 77, 281, 1073, 4193, 16577, 65921}
	case "sedgewick2":
		gaps = []int{1, 5, 19, 41, 109, 209, 505, 929, 2161, 3905, 8929, 16001, 36289, 64769}
	}
	var i int
	for i = 0; i < len(gaps) && gaps[i] <= size; i++ {
	}
	gaps = gaps[:i]
	reverse(gaps)
	return gaps
}

// Shell Gap Sequence
// ex: n/2, n/4, n/8...
func shellGaps(size int) []int {
	gaps := []int{}
	for i := size; i > 0; i /= 2 {
		gaps = append(gaps, i)
	}
	return gaps
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
