package dijkstra

import (
	"fmt"
	"log"
)

// denseGraph 邻接矩阵-稠密图
// vertexNum 顶点数量
// adjMatrix 邻接矩阵存储结构
type denseGraph struct {
	vertexNum int
	adjMatrix [][]float64
}

// NewDenseGraph 新建空稠密图，vertexNum表示该图的顶点数
func NewDenseGraph(vertexNum int) *denseGraph {
	adjMatrix := make([][]float64, vertexNum)
	for i := 0; i < vertexNum; i++ {
		adjMatrix[i] = make([]float64, vertexNum)
		for j := 0; j < vertexNum; j++ {
			adjMatrix[i][j] = Max
		}
	}
	return &denseGraph{vertexNum: vertexNum, adjMatrix: adjMatrix}
}

func (dg *denseGraph) VertexNum() int {
	return dg.vertexNum
}

func (dg *denseGraph) SetVertexNum(vertexNum int) {
	dg.vertexNum = vertexNum
}

func (dg *denseGraph) AddEdge(sv, ev int, weight float64) {
	if sv >= dg.vertexNum || ev >= dg.vertexNum {
		log.Printf("顶点超过限制数量")
		return
	}
	dg.adjMatrix[sv][ev] = weight
}

func (dg *denseGraph) AddDoubleEdge(sv, ev int, weight float64) {
	dg.AddEdge(sv, ev, weight)
	dg.AddEdge(ev, sv, weight)
}

func (dg denseGraph) getWeight(sv, ev int) (float64, bool) {
	if dg.adjMatrix[sv][ev] != Max {
		return dg.adjMatrix[sv][ev], true
	}
	return Max, false
}

// String 输出图
func (dg denseGraph) String() string {
	for i := 0; i < dg.vertexNum; i++ {
		for j := 0; j < dg.vertexNum; j++ {
			if dg.adjMatrix[i][j] != Max {
				fmt.Print(dg.adjMatrix[i][j], " ")
			} else {
				fmt.Print("MAX ")
			}
		}
		fmt.Println()
	}
	return ""
}
