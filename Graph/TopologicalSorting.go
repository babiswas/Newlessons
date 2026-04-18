package arr

import (
	"container/list"
	"fmt"
)

type Topsort struct {
	vertex int
	adj    map[int]*list.List
}

func (t *Topsort) topsort_init(vertex int) {
	t.vertex = vertex
	t.adj = make(map[int]*list.List)
	for i := range t.vertex {
		t.adj[i] = list.New()
	}
}

func (t *Topsort) add_edges(i, j int) {
	_, ok := t.adj[i]
	if !ok {
		t.adj[i] = list.New()
	}
	t.adj[i].PushBack(j)
}

func (t *Topsort) topsort() {
	visited := make([]bool, t.vertex)
	for i := 0; i < t.vertex; i++ {
		visited[i] = false
	}
	stack := list.New()
	for i := range t.vertex {
		if !visited[i] {
			t.topsort_util(i, visited, stack)
		}
	}

	for elm := stack.Back(); elm != nil; elm = elm.Prev() {
		fmt.Println(elm.Value.(int))
	}
}

func (t *Topsort) topsort_util(i int, visited []bool, stack *list.List) {
	visited[i] = true
	temp_list := t.adj[i]
	for elm := temp_list.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if !visited[index] {
			t.topsort_util(index, visited, stack)
		}
	}
	stack.PushBack(i)
}

func TopologicalSorting() {
	fmt.Println("Topological sorting:")
	topsort := Topsort{}
	fmt.Println("Constructing the graph:")
	topsort.topsort_init(6)
	topsort.add_edges(5, 0)
	topsort.add_edges(4, 0)
	topsort.add_edges(5, 2)
	topsort.add_edges(2, 3)
	topsort.add_edges(3, 1)
	topsort.add_edges(4, 1)
	fmt.Println("Performing topological sorting:")
	topsort.topsort()
}
