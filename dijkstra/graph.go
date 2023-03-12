package dijkstra

// graph 图接口，只要是实现此接口的结构体都可执行dijkstra算法
// 后续可对graph接口进行扩展支持其他导航算法
type graph interface {
	//VertexNum 获取顶点数量
	VertexNum() int

	//SetVertexNum 设置顶点数量
	SetVertexNum(vertexNum int)

	//AddEdge 添加边 sv表示起始顶点 ev表示结束顶点
	AddEdge(sv, ev int, weight float64)

	//AddDoubleEdge 添加两条边，用于无向图添加
	//weight表示由sv与ev构成的边的权重(导航中为两点间距离)
	AddDoubleEdge(sv, ev int, weight float64)

	//getWeight 获取边的权重(点间距离)，若不存在权重正无穷
	getWeight(sv, ev int) (float64, bool)
}
