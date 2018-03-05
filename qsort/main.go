package main

import (
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var a []int
	size := 30
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(3))
	}
	qsort(a, 0, len(a)-1)
}

func qsort(a []int, lo int, hi int) {
	logrus.Info("--------------------------------------------------------------------------------")
	logrus.Infof("%v qsort lo=%d, hi=%d", a, lo, hi)
	if lo < hi {
		// p := hoarePartition(a, lo, hi)
		p := cormenPartition(a, lo, hi)
		qsort(a, lo, p-1)
		qsort(a, p+1, hi)
	}
}

func cormenPartition(a []int, lo int, hi int) int {
	pivot := a[hi]
	logrus.Infof("pivot=%d", pivot)
	i := lo - 1
	for j := lo; j <= hi-1; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	if a[hi] < a[i+1] {
		a[i+1], a[hi] = a[hi], a[i+1]
	}
	return i + 1
}

func hoarePartition(a []int, lo int, hi int) int {
	// pivot
	// p := rand.Intn(len(a))
	p := 8
	pivot := a[p]
	spew.Dump(p)
	logrus.Infof("%v", a)
	i := lo - 1
	j := hi + 1
	for {
		logrus.Info("-----------------------------")
		spew.Dump(i, j)
		logrus.Infof("a[%d]=%2d j[%d]=%2d p[%d]=%2d", i, a[i], j, a[j], p, a[p])
		for a[i] < pivot {
			logrus.Info("i++")
			i++
		}

		for a[j] >= pivot {
			logrus.Info("j--")
			j--
		}

		if i >= j {
			return j
		}

		logrus.Infof("swap %d <> %d", a[i], a[j])
		a[i], a[j] = a[j], a[i]
		logrus.Infof("%v", a)
	}
}
