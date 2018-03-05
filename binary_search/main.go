package main

import (
	"fmt"
	"math/rand"

	"github.com/Sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	var a []int
	size := 100000
	logrus.Debugf("generating list size: %d", size)
	for i := 0; i < size; i++ {
		a = append(a, i+1)
	}
	x := rand.Intn(size)
	logrus.Debugf("searching...")
	idx, found := binarySearch(a, x, 0, len(a)-1)
	if !found {
		fmt.Printf("%d not found\n", x)
		return
	}
	fmt.Printf("%d found at %d\n", x, idx)
}

var level = 0

func binarySearch(a []int, x, l, h int) (int, bool) {
	level++

	logrus.Debugf("recursing stack: %d, n=%d, i=%d, j=%d", level, x, l, h)

	if l > h {
		return 0, false
	}

	m := (l + h) / 2
	e := a[m]

	if x == e {
		return m, true
	} else if x < e {
		return binarySearch(a, x, l, m-1)
	}
	return binarySearch(a, x, m+1, h)

}
