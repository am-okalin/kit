package dijkstra

import (
	"testing"
)

func createGraph() Graph {
	graph := NewDenseGraph(6)
	//Graph := NewSparseGraph(6)
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 3, 4)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(2, 1, 2)
	graph.AddEdge(2, 3, 9)
	graph.AddEdge(2, 5, 3)
	graph.AddEdge(3, 5, 1)
	return graph
}

func Test_dijkstra0(t *testing.T) {
	graph := createGraph()

	prev, dist := Dijkstra(graph, 0)
	t.Log(prev)
	t.Log(dist)

	prev, dist = Dijkstra2(graph, 1)
	t.Log(prev)
	t.Log(dist)

	var res []int
	GetPrev(5, prev, &res)
	t.Log(res)
}

func Test_dijkstra1(t *testing.T) {
	graph := createGraph()

	prev, dist := Dijkstra(graph, 1)
	t.Log(prev)
	t.Log(dist)

	prev, dist = Dijkstra2(graph, 1)
	t.Log(prev)
	t.Log(dist)
}

func Test_dijkstra2(t *testing.T) {
	graph := createGraph()

	prev, dist := Dijkstra(graph, 2)
	t.Log(prev)
	t.Log(dist)

	prev, dist = Dijkstra2(graph, 2)
	t.Log(prev)
	t.Log(dist)
}
