package arr

import (
	"container/list"
	"fmt"
)

type BFS struct {
	vertex int
	adj    map[int]*list.List
}

func (g *BFS) BFSinit(vertex int) {
	g.vertex = vertex
	g.adj = make(map[int]*list.List)
	for i := 0; i < g.vertex; i++ {
		g.adj[i] = list.New()
	}
}

func (g *BFS) AddEdges(i, j int) {
	_, ok := g.adj[i]
	if !ok {
		g.adj[i] = list.New()
	}
	g.adj[i].PushBack(j)
}

func (g *BFS) bfsTraversal(i int) {
	visited := make([]bool, g.vertex)
	for i := 0; i < g.vertex; i++ {
		visited[i] = false
	}
	queue := list.New()
	queue.PushBack(i)
	visited[i] = true
	for queue.Len() != 0 {
		elm := queue.Front()
		index := elm.Value.(int)
		fmt.Println(index)
		temp_list := g.adj[index]
		queue.Remove(elm)
		for em := temp_list.Front(); em != nil; em = em.Next() {
			val := em.Value.(int)
			if !visited[val] {
				queue.PushBack(val)
				visited[val] = true
			}
		}
	}
}

func BFSTraversal() {
	fmt.Println("BFS traversal:")
	bfs := BFS{}
	fmt.Println("Constructing the graph:")
	bfs.BFSinit(4)
	bfs.AddEdges(0, 2)
	bfs.AddEdges(2, 0)
	bfs.AddEdges(0, 1)
	bfs.AddEdges(1, 2)
	bfs.AddEdges(2, 3)
	bfs.AddEdges(3, 3)
	fmt.Println("Traversing the graph:")
	bfs.bfsTraversal(2)
}
