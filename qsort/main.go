package main

import (
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	var a []int
	size := 10
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(49))
	}
	qsort(a, 0, len(a)-1)
}

func qsort(a []int, lo int, hi int) {
	logrus.Info("-----------------------------------------------------------------------------------------------")
	// logrus.Infof("lo=%d, hi=%d\n", a, lo, hi)
	if lo < hi {
		// p := hoarePartition(a, lo, hi)
		p := cormenPartition(a, lo, hi)
		qsort(a, lo, p-1)
		qsort(a, p+1, hi)
	}
}

func cormenPartition(a []int, lo int, hi int) int {
	pivot := a[hi]
	logrus.Infof("pivot=%d lo=%d hi=%d", pivot, lo, hi)
	i := lo - 1
	for j := lo; j <= hi-1; j++ {
		// logrus.Debugf("lo=%d hi=%d i=%d j=%d", lo, hi, i, j)
		if a[j] < pivot {
			// logrus.Debugf("swapping")
			i++
			a[i], a[j] = a[j], a[i]
		}
		logrus.Debugf("%v lo=%d hi=%d i=%d j=%d pivot=%d", a, lo, hi, i, j, pivot)
	}
	if a[hi] < a[i+1] {
		a[i+1], a[hi] = a[hi], a[i+1]
		logrus.Debugf("%v", a)
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
