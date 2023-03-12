package dijkstra

import (
	"container/list"
	"fmt"
	"log"
)

// sparseGraph 邻接表-稀疏图
// vertexNum 顶点数量
// adjTable 邻接表是一个链表结构
type sparseGraph struct {
	vertexNum int
	adjTable  []*list.List
}

// NewSparseGraph 新建空稀疏图，vertexNum表示该图的顶点数
func NewSparseGraph(vertexNum int) *sparseGraph {
	adjTable := make([]*list.List, vertexNum)
	for i := 0; i < vertexNum; i++ {
		adjTable[i] = list.New()
	}

	return &sparseGraph{vertexNum: vertexNum, adjTable: adjTable}
}

func (sg *sparseGraph) VertexNum() int {
	return sg.vertexNum
}

func (sg *sparseGraph) SetVertexNum(vertexNum int) {
	sg.vertexNum = vertexNum
}

func (sg *sparseGraph) AddEdge(sv, ev int, weight float64) {
	if sv >= sg.vertexNum || ev >= sg.vertexNum {
		log.Printf("顶点超过限制数量")
		return
	}
	sg.adjTable[sv].PushBack(&Elem{ev: ev, weight: weight})
}

func (sg *sparseGraph) AddDoubleEdge(sv, ev int, weight float64) {
	sg.AddEdge(sv, ev, weight)
	sg.AddEdge(ev, sv, weight)
}

func (sg sparseGraph) getWeight(sv, ev int) (float64, bool) {
	for li := sg.adjTable[sv].Front(); li != nil; li = li.Next() {
		elem := li.Value.(*Elem)
		if elem.ev == ev {
			return elem.weight, true
		}
	}
	return Max, false
}

// String 输出图
func (sg sparseGraph) String() string {
	for i := 0; i < sg.vertexNum; i++ {
		fmt.Print(i, ": ")
		for li := sg.adjTable[i].Front(); li != nil; li = li.Next() {
			elem := li.Value.(*Elem)
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
	return ""
}
