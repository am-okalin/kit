package dijkstra

import (
	"container/heap"
	"testing"
)

// TestHeap 测试堆的添加/取出/排序元素是否正确
func TestHeap(t *testing.T) {
	h := &Heap{}
	heap.Push(h, &Elem{3, 5})
	heap.Push(h, &Elem{2, 3})
	heap.Push(h, &Elem{4, 4})
	heap.Push(h, &Elem{1, 1})
	hl := h.Len()
	t.Log(hl)
	for i := 0; i < hl; i++ {
		t.Log(heap.Pop(h))
	}
	t.Log(h.Len())
}
