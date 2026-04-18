package arr

import (
	"container/list"
	"fmt"
)

type TransitiveClosure struct {
	vertex int
	adj    map[int]*list.List
}

func (t *TransitiveClosure) graph_init(vertex int) {
	t.vertex = vertex
	t.adj = make(map[int]*list.List)
	for i := 0; i < t.vertex; i++ {
		t.adj[i] = list.New()
	}
}

func (t *TransitiveClosure) add_edges(i, j int) {
	_, ok := t.adj[i]
	if !ok {
		t.adj[i] = list.New()
	}
	t.adj[i].PushBack(j)
}

func (t *TransitiveClosure) transitive_closure_util(i, j int, arr [][]int) {
	arr[i][j] = 1
	tmp_list := t.adj[j]
	for el := tmp_list.Front(); el != nil; el = el.Next() {
		index := el.Value.(int)
		if arr[i][index] != 1 {
			t.transitive_closure_util(i, index, arr)
		}
	}
}

func (t *TransitiveClosure) transitive_closure() {
	arr := make([][]int, t.vertex)
	for i := 0; i < t.vertex; i++ {
		arr[i] = make([]int, t.vertex)
	}
	for i := 0; i < t.vertex; i++ {
		t.transitive_closure_util(i, i, arr)
	}
	for i := 0; i < t.vertex; i++ {
		fmt.Println(arr[i])
	}
}

func Transitive_Closure() {
	fmt.Println("TRansitive closure of a graph:")
	tclosure := TransitiveClosure{}
	tclosure.graph_init(4)
	tclosure.add_edges(2, 0)
	tclosure.add_edges(0, 2)
	tclosure.add_edges(2, 3)
	tclosure.add_edges(3, 3)
	tclosure.add_edges(0, 1)
	tclosure.add_edges(1, 2)
	fmt.Println("Transitive closure of the graph:")
	tclosure.transitive_closure()
}
