package heap

import (
	//"fmt"
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)
}

type BinaryHeap struct {
	nodes []int
}

// left return the index of the left child for the given node index
func (b *BinaryHeap) left(i int) (int, bool) {
	idx := 2*i + 1
	return idx, idx < len(b.nodes)
}

// right returns the index of the right child for the given node index
func (b *BinaryHeap) right(i int) (int, bool) {
	idx := 2*i + 2
	return idx, idx < len(b.nodes)
}

// parent returns the index of the parent for the given node index
func (b *BinaryHeap) parent(i int) int {
	return (i - 1) / 2
}

// balance ensures that the given index is properly places in the head
func (b *BinaryHeap) bubble(i int) int {
	p := b.parent(i)
	for i > 0 && b.nodes[i] < b.nodes[p] {
		log.WithFields(log.Fields{"index": i, "value": b.nodes[i], "nodes": b.nodes}).Debug("Bubbling")
		v := b.nodes[i]
		b.nodes[i] = b.nodes[p]
		b.nodes[p] = v
		i = p
		p = b.parent(i)
	}
	return i
}

// Add inserts the given value into the heap.
func (b *BinaryHeap) Add(v int) int {
	log.WithFields(log.Fields{"value": v}).Info("Adding value")
	b.nodes = append(b.nodes, v)
	return b.bubble(len(b.nodes) - 1)
}

// Remove returns the highest priority item from the head,
func (b *BinaryHeap) Remove() int {

	if len(b.nodes) <= 0 {
		log.Warn("Remove: Head is empty")
		return -1
	}

	smallest := b.nodes[0]

	log.WithFields(log.Fields{"value": smallest}).Info("Removing Value")

	j := len(b.nodes) - 1
	b.nodes[0] = b.nodes[j]
	b.nodes = b.nodes[:j]
	if j > 0 {
		b.trickle(0)
	}
	return smallest
}

// trickle rebalances the heap after a node removal
func (b *BinaryHeap) trickle(i int) {
	log.WithFields(log.Fields{"index": i, "value": b.nodes[i], "nodes": b.nodes}).Debug("Trickling")
	lc, lcExists := b.left(i)
	rc, rcExists := b.right(i)
	j := -1

	if rcExists && b.nodes[rc] < b.nodes[i] {
		if b.nodes[lc] < b.nodes[rc] {
			j = lc
		} else {
			j = rc
		}
	} else {
		if lcExists && b.nodes[lc] < b.nodes[i] {
			j = lc
		}
	}

	if j < 0 {
		return
	}

	// swap the child with the parent
	tmp := b.nodes[i]
	b.nodes[i] = b.nodes[j]
	b.nodes[j] = tmp

	b.trickle(j)
}
