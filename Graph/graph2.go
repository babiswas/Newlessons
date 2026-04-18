package arr

import (
	"container/list"
	"fmt"
)

type DFS struct {
	vertex int
	adj    map[int]*list.List
}

func (d *DFS) dfsinit(vertex int) {
	d.vertex = vertex
	d.adj = make(map[int]*list.List)
	for i := range d.vertex {
		d.adj[i] = list.New()
	}
}

func (d *DFS) add_edges(i, j int) {
	_, ok := d.adj[i]
	if !ok {
		d.adj[i] = list.New()
	}
	d.adj[i].PushBack(j)
}

func (d *DFS) dfs() {
	visited := make([]bool, d.vertex)
	for i := range d.vertex {
		visited[i] = false
	}
	for j := range d.vertex {
		if !visited[j] {
			d.dfs_util(j, visited)
		}
	}
}

func (d *DFS) dfs_util(j int, visited []bool) {
	visited[j] = true
	fmt.Println(j)
	temp_list := d.adj[j]
	for elm := temp_list.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if !visited[index] {
			d.dfs_util(index, visited)
		}
	}
}

func DFSTraversal() {
	fmt.Println("performing dfs traversal")
	dfsobj := DFS{}
	dfsobj.dfsinit(4)
	fmt.Println("Constructing graph:")
	dfsobj.add_edges(0, 2)
	dfsobj.add_edges(2, 0)
	dfsobj.add_edges(0, 1)
	dfsobj.add_edges(1, 2)
	dfsobj.add_edges(2, 3)
	dfsobj.add_edges(3, 3)
	fmt.Println("Performing dfs traversal:")
	dfsobj.dfs()
}
