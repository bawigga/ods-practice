package heap

type Heap struct {
	nodes []int
}

// NewHeap returns a newly initialized Heap
func NewHeap() Heap {
	return Heap{nodes: []int{}}
}

// Insert the element into the Heap and returns its index
func (h *Heap) Insert(n int) int {
	h.nodes = append(h.nodes, n)

	i := len(h.nodes) - 1
	for {
		p := parent(i)
		if n >= h.nodes[p] {
			break
		}
		h.nodes[i], h.nodes[p] = h.nodes[p], h.nodes[i]
		i = p
	}

	return i
}

func (h *Heap) Remove() int {

	r := h.nodes[0]

	// reduce nodes by 1
	li := len(h.nodes) - 1
	h.nodes[0] = h.nodes[li]
	h.nodes = h.nodes[:li]
	li--

	i := 0
	for {
		l := left(i)
		r := right(i)

		if l <= li {
			if r <= li {
				if h.nodes[l] <= h.nodes[r] {
					h.nodes[i], h.nodes[l] = h.nodes[l], h.nodes[i]
					i = l
				} else {
					h.nodes[i], h.nodes[r] = h.nodes[r], h.nodes[i]
					i = r
				}
			} else {
				h.nodes[i], h.nodes[l] = h.nodes[l], h.nodes[i]
				i = l
			}
		} else {
			break
		}
	}

	return r
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}
