package heap

import (
	_ "fmt"
)

type BinaryHeap struct {
	nodes []int
}

func (b *BinaryHeap) left(i int) int {
	return 2*i + 1
}

func (b *BinaryHeap) right(i int) int {
	return 2*i + 2
}

func (b *BinaryHeap) parent(i int) int {
	return (i - 1) / 2
}

func (b *BinaryHeap) balance(i int) int {
	p := b.parent(i)
	for i > 0 && b.nodes[i] < b.nodes[p] {
		v := b.nodes[i]
		b.nodes[i] = b.nodes[p]
		b.nodes[p] = v
		i = p
		p = b.parent(i)
	}
	return i
}

func (b *BinaryHeap) Add(v int) int {
	b.nodes = append(b.nodes, v)
	return b.balance(len(b.nodes) - 1)
}
