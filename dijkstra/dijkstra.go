package dijkstra

import (
	"container/heap"
	"math"
)

// initPrev 初始化`起始顶点(sv)到任意一点的`路径
// n表示prev长度，即顶点数
func initPrev(n int) []int {
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = -1
	}
	return prev
}

// initDist 初始化`起始顶点(sv)到任意一点的`距离
// sv表示起始顶点
func initDist(n, sv int) []float64 {
	dist := make([]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxFloat64
	}
	//sv到自己的距离为0
	dist[sv] = 0
	return dist
}

// initQue 初始化优先级队列，用堆结构作为优先级队列
func initQue(sv int) *Heap {
	que := &Heap{}
	heap.Push(que, &Elem{sv, 0})
	return que
}

// getShortest 从 dist[e] (e∈s) 中获取最短路径的的顶点ID
// g表示实现graph接口的类实例
// s表示所有点的集合，true表示已找到最短路径
// dist表示起始顶点(sv)到任意一点的距离
func getShortest(g Graph, s map[int]bool, dist []float64) int {
	mink := -1
	min := math.MaxFloat64
	for j := 0; j < g.VertexNum(); j++ {
		if s[j] == false && dist[j] < min {
			min = dist[j]
			mink = j
		}
	}
	return mink
}

// Dijkstra 找到sv到每个点的最短路径：每次只走距离原点最近的点，通过此路径优化其他路径
// 此算法可以在找到ev后立即停下来,但这里只提供了找出所有点的最短路径方法(是想知道在图过大时找最远点是否会超时，如果超时建议选择别的算法)
func Dijkstra(g Graph, sv int) ([]int, []float64) {
	//s是所有点的集合，true表示已找到最短路径
	n := g.VertexNum()
	s := map[int]bool{}
	prev := initPrev(n)
	dist := initDist(n, sv)
	//找出sv至所有点的最近点
	for i := 0; i < n-1; i++ {
		mink := getShortest(g, s, dist)
		if mink == -1 {
			break
		}
		s[mink] = true

		//对新扩展的路线进行`松弛`
		//松弛:设当前最近点为iv，若`sv->iv->其他顶点`大于`sv->其他顶点`则对`prev`和`dist`进行修改
		for j := 0; j < n; j++ {
			if mink == j {
				continue
			}
			weight, hasRoute := g.getWeight(mink, j)
			if hasRoute {
				w := dist[mink] + weight
				if w < dist[j] {
					dist[j] = w
					prev[j] = mink
				}
			}
		}
	}
	return prev, dist
}

// Dijkstra2 为dijkstra的改进算法，使用了优先级队列进行了优化
func Dijkstra2(g Graph, sv int) ([]int, []float64) {
	//s是所有点的集合，true表示已找到最短路径
	s := map[int]bool{}
	n := g.VertexNum()
	prev := initPrev(n)
	dist := initDist(n, sv)
	que := initQue(sv)
	//找出sv至所有点的最近点
	for que.Len() != 0 {
		mink := heap.Pop(que).(*Elem).ev

		s[mink] = true

		//对新扩展的路线进行松弛
		for j := 0; j < n; j++ {
			if mink == j {
				continue
			}
			weight, hasRoute := g.getWeight(mink, j)
			if hasRoute {
				w := dist[mink] + weight
				if w < dist[j] {
					dist[j] = w
					prev[j] = mink
					//将 w, j 加入优先级队列
					heap.Push(que, &Elem{j, w})
				}
			}
		}
	}
	return prev, dist
}

// GetPrev (尾递归) 获取结果集的标准最短路径表示
func GetPrev(ev int, prev []int, r *[]int) {
	if ev == -1 {
		return
	}
	GetPrev(prev[ev], prev, r)
	*r = append(*r, ev)
}
