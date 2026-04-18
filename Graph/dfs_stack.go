package arr

import (
	"container/list"
	"fmt"
)

type DFSStack struct {
	vertex int
	adj    map[int]*list.List
}

func (d *DFSStack) dfsstack_init(vertex int) {
	d.vertex = vertex
	d.adj = make(map[int]*list.List)
	for i := range d.vertex {
		d.adj[i] = list.New()
	}
}

func (d *DFSStack) add_edges(i, j int) {
	_, ok := d.adj[i]
	if !ok {
		d.adj[i] = list.New()
	}
	d.adj[i].PushBack(j)
}

func (d *DFSStack) dfs_stack_util(i int) {
	visited := make([]bool, d.vertex)
	for i := range d.vertex {
		visited[i] = false
	}
	stack := list.New()
	stack.PushBack(i)
	for stack.Len() > 0 {
		elm := stack.Back()
		index := elm.Value.(int)
		if !visited[index] {
			fmt.Println(index)
			visited[index] = true
		}
		stack.Remove(elm)
		temp_list := d.adj[index]
		for el := temp_list.Front(); el != nil; el = el.Next() {
			num := el.Value.(int)
			if !visited[num] {
				stack.PushBack(num)
			}
		}
	}
}

func DFS_Stack() {
	fmt.Println("DFS traversal using stack:")
	dfs_stack := DFSStack{}
	dfs_stack.dfsstack_init(4)

	fmt.Println("Constructing the graph:")
	dfs_stack.add_edges(0, 2)
	dfs_stack.add_edges(2, 0)
	dfs_stack.add_edges(0, 1)
	dfs_stack.add_edges(1, 2)
	dfs_stack.add_edges(2, 3)
	dfs_stack.add_edges(3, 3)

	fmt.Println("DFS traversal:")
	dfs_stack.dfs_stack_util(2)
}
