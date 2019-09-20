package commonds

import (
	"errors"
	"fmt"
)

/**
图数据结构，这里使用邻接表来存储, 所有的顶点和邻接表存储在map中,key为顶点,value为邻接表
directed标识是有向图还是无向图: true 有向图 false 无向图
orderkeys 按照插入顺序排序的顶点数组,因为range遍历map是无序的,有时候会有有序的需求
 */
type Graph struct {
	directed bool
	orderkeys []interface{}
	vertices map[interface{}]*LinkedList
}

/**
创建无向图 Graph
 */
func NewNDGraph() *Graph {
	return &Graph{directed: false, vertices: make(map[interface{}]*LinkedList)}
}

/**
创建有向图 Graph
 */
func NewDGraph() *Graph {
	return &Graph{directed: true, vertices: make(map[interface{}]*LinkedList)}
}

/**
插入顶点
vetice 顶点数据
 */
func (graph *Graph) AddVertice(vetice interface{}) error {
	_, ok := graph.vertices[vetice]
	if ok {
		return errors.New("vetice exists")
	} else {
		graph.vertices[vetice] = NewLinkedList()
		graph.orderkeys = append(graph.orderkeys, vetice)
	}
	return nil
}

/**
插入边
fvetice: 边的起始顶点
tvetice: 边的终止顶点
 */
func (graph *Graph) AddEdge(fvetice interface{}, tvetice interface{}) error {
	fnt, ok := graph.vertices[fvetice]
	if !ok {
		if err := graph.AddVertice(fvetice); err != nil {
			return err
		}
	}
	fnt.Insert(tvetice)
	// 无向图的边是等价的，所以需要双向插入
	if !graph.directed {
		tnt, ok := graph.vertices[tvetice]
		if !ok {
			if err := graph.AddVertice(tvetice); err != nil {
				return err
			}
		}
		tnt.Insert(fvetice)
	}
	return nil
}

/**
深度优先搜索
 */
func (graph *Graph) DeepFirstSearch(vetice interface{}, optr func(data interface{}) bool) {

}

/**
显示图结构
 */
func (graph *Graph) Show() {
	for _, key := range graph.orderkeys {
		fmt.Print(key, "--->")
		nt, ok := graph.vertices[key]
		if ok {
			nt.Operator(func(data interface{}) bool {
				fmt.Print(data, " ")
				return true
			})
		}
		fmt.Println()
	}
}