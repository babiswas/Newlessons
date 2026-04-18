package arr

import (
	"container/list"
	"fmt"
)

type IsCyclic struct {
	vertex int
	adj    map[int]*list.List
}

func (c *IsCyclic) cyclic_init(vertex int) {
	c.vertex = vertex
	c.adj = make(map[int]*list.List)
	for i := 0; i < c.vertex; i++ {
		c.adj[i] = list.New()
	}
}

func (c *IsCyclic) add_edges(i, j int) {
	_, ok := c.adj[i]
	if !ok {
		c.adj[i] = list.New()
	}
	c.adj[i].PushBack(j)
}

func (c *IsCyclic) is_cyclic() bool {
	visited := make([]bool, c.vertex)
	recstack := make([]bool, c.vertex)
	for i := 0; i < c.vertex; i++ {
		visited[i] = false
		recstack[i] = false
	}
	for i := 0; i < c.vertex; i++ {
		if !visited[i] {
			if c.is_cyclic_util(i, visited, recstack) {
				return true
			}
		}
	}
	return false
}

func (c *IsCyclic) is_cyclic_util(u int, visited []bool, recstack []bool) bool {
	visited[u] = true
	recstack[u] = true
	temp_list := c.adj[u]
	for el := temp_list.Front(); el != nil; el = el.Next() {
		index := el.Value.(int)
		if !visited[index] {
			if c.is_cyclic_util(index, visited, recstack) {
				return true
			}
		} else if recstack[index] {
			return true
		}
	}
	recstack[u] = false
	return false
}

func VerifyCycle() {
	fmt.Println("Verify if there is a cycle in the graph:")
	iscyclic := IsCyclic{}
	iscyclic.cyclic_init(4)
	iscyclic.add_edges(0, 2)
	iscyclic.add_edges(2, 0)
	iscyclic.add_edges(2, 3)
	iscyclic.add_edges(3, 3)
	iscyclic.add_edges(0, 1)
	iscyclic.add_edges(1, 2)
	fmt.Println("Detect cycle:")
	if iscyclic.is_cyclic() {
		fmt.Println("There is a cycle in the graph:")
	} else {
		fmt.Println("There is no cycle in the graph")
	}
}
