package dijkstra

// Elem 导航中优先级队列中的堆元素
// ev 结束顶点(end vertex)
// weight 边的权重(也就是顶点与顶点之间的距离)
type Elem struct {
	ev     int
	weight float64
}

// Heap 用堆实现优先级队列
type Heap []*Elem

// Less 构造的是小顶堆，大顶堆需要修改下面的`<`符号
func (h *Heap) Less(i, j int) bool {
	return (*h)[i].weight < (*h)[j].weight
}

// Swap 交换堆中两结点指
func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Len 获取堆的长度
func (h *Heap) Len() int {
	return len(*h)
}

// Pop 从堆中取出元素
func (h *Heap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

// Push 往堆中添加元素
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*Elem))
}
