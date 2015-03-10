package heap

import (
	"fmt"
	"testing"
)

func TestBinaryHeapNew(t *testing.T) {
	//heap := BinaryHeap{}
}

func TestAdd(t *testing.T) {
	heap := BinaryHeap{
		nodes: []int{4, 9, 8, 17, 26, 50, 16, 19, 69, 32, 93, 55},
	}
	index := heap.Add(24)
	index = heap.Add(96)
	index = heap.Add(62)
	index = heap.Add(2)
	index = heap.Add(8)
	if index != 2 {
		t.Errorf("Expected index %d, got %d", 2, index)
	}
}

func TestRemove(t *testing.T) {
	heap := BinaryHeap{
		nodes: []int{4, 9, 6, 17, 26, 8, 16, 19, 69, 32, 93, 55, 50},
	}

	nodes := []int{}
	node := heap.Remove()
	for node > 0 {
		nodes = append(nodes, node)
		node = heap.Remove()
	}
	fmt.Printf("nodes", nodes)
}
