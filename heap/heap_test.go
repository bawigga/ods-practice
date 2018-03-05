package heap

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	heap := Heap{
		nodes: []int{},
	}
	index := heap.Insert(10)
	if index != 0 {
		t.Errorf("Expected index %d, got %d", 0, index)
	}

	index = heap.Insert(30)
	if index != 1 {
		t.Errorf("Expected index %d, got %d", 1, index)
	}

	index = heap.Insert(20)
	index = heap.Insert(15)
	if index != 1 {
		t.Errorf("Expected index %d, got %d", 1, index)
	}

	index = heap.Insert(1)
	if index != 0 {
		t.Errorf("Expected index %d, got %d", 0, index)
	}
}

func TestRemove(t *testing.T) {
	heap := Heap{
		nodes: []int{4, 9, 6, 17, 26, 8, 16},
	}

	items := []int{}
	for {
		items = append(items, heap.Remove())
		if len(heap.nodes) == 0 {
			break
		}
	}

	sorted := []int{4, 6, 8, 9, 16, 17, 26}
	if !reflect.DeepEqual(items, sorted) {
		t.Error("did not remove items in order")
	}
}

// func TestHeapify(t *testing.T) {
// }
