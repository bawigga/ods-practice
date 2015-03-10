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
	index := heap.Add(6)
	if index != 2 {
		t.Errorf("Expected index %d, got %d", 2, index)
	}
	//fmt.Printf("nodes: %v", heap)
}
