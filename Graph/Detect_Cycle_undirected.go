package arr

import (
	"container/list"
	"fmt"
)

type UndirectedCyclic struct {
	vertex int
	adj    map[int]*list.List
}

func (u *UndirectedCyclic) graph_init(vertex int) {
	u.vertex = vertex
	u.adj = make(map[int]*list.List)
	for i := 0; i < u.vertex; i++ {
		u.adj[i] = list.New()
	}
}

func (u *UndirectedCyclic) add_edges(i, j int) {
	_, ok := u.adj[i]
	if !ok {
		u.adj[i] = list.New()
	}
	u.adj[i].PushBack(j)
	_, ok = u.adj[j]
	if !ok {
		u.adj[j] = list.New()
	}
	u.adj[j].PushBack(i)
}

func (u *UndirectedCyclic) isCyclic() bool {
	visited := make([]bool, u.vertex)
	for i := 0; i < u.vertex; i++ {
		visited[i] = false
	}
	for i := 0; i < u.vertex; i++ {
		if !visited[i] {
			if u.iscyclic_util(i, visited, -1) {
				return true
			}
		}
	}
	return false
}

func (u *UndirectedCyclic) iscyclic_util(i int, visited []bool, parent int) bool {
	visited[i] = true
	temp_list := u.adj[i]
	for el := temp_list.Front(); el != nil; el = el.Next() {
		index := el.Value.(int)
		if !visited[index] {
			if u.iscyclic_util(index, visited, i) {
				return true
			}
		} else if index != parent {
			return true
		}
	}
	return false
}

func CyclicUndirected() {
	fmt.Println("Verify if the undirected graph is cyclic.")
	uc := UndirectedCyclic{}
	uc.graph_init(4)
	fmt.Println("Constructing the graph:")
	uc.add_edges(0, 2)
	uc.add_edges(2, 3)
	uc.add_edges(0, 1)
	fmt.Println("Detecting cycle:")
	if uc.isCyclic() {
		fmt.Println("There is a cycle in the graph:")
	} else {
		fmt.Println("There is no cycle in the graph.")
	}
}
