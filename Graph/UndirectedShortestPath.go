package arr

import (
	"container/list"
	"fmt"
)

type ShortestPath struct {
	vertex int
	adj    map[int]*list.List
}

func (s *ShortestPath) shortestPathInit(vertex int) {
	s.vertex = vertex
	s.adj = make(map[int]*list.List)
	for i := 0; i < s.vertex; i++ {
		s.adj[i] = list.New()
	}
}

func (s *ShortestPath) add_edges(i, j int) {
	_, ok := s.adj[i]
	if !ok {
		s.adj[i] = list.New()
	}
	s.adj[i].PushBack(j)
	_, ok = s.adj[j]
	if !ok {
		s.adj[j] = list.New()
	}
	s.adj[j].PushBack(i)
}

func (s *ShortestPath) bfs(src int) []int {
	prev := make([]int, s.vertex)
	for i := 0; i < s.vertex; i++ {
		prev[i] = -1
	}
	visited := make([]bool, s.vertex)
	for i := 0; i < s.vertex; i++ {
		visited[i] = false
	}
	queue := list.New()
	queue.PushBack(src)
	visited[src] = true
	for queue.Len() > 0 {
		elm := queue.Front()
		index := elm.Value.(int)
		queue.Remove(elm)
		temp_list := s.adj[index]
		for el := temp_list.Front(); el != nil; el = el.Next() {
			num := el.Value.(int)
			if !visited[num] {
				visited[num] = true
				queue.PushBack(num)
				prev[num] = index
			}
		}
	}
	fmt.Println(prev)
	return prev
}

func (s *ShortestPath) ShortestPath(source, target int) {
	prev := s.bfs(source)
	path := make([]int, 0)
	path = append(path, target)
	index := prev[target]
	for index != -1 {
		path = append(path, index)
		index = prev[index]
	}
	reverse := func(arr []int, start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start += 1
			end -= 1
		}
	}
	reverse(path, 0, len(path)-1)
	if path[0] == source {
		fmt.Println("There is a path from source,", source, "target,", target)
		fmt.Println(path)
	} else {
		fmt.Println("Shortest path don't exist.")
	}
}

func ShortestPathUndirected() {
	fmt.Println("Find shortest path in undirected graph:")
	shortest_path := ShortestPath{}
	shortest_path.shortestPathInit(8)
	shortest_path.add_edges(1, 2)
	shortest_path.add_edges(1, 0)
	shortest_path.add_edges(0, 3)
	shortest_path.add_edges(3, 7)
	shortest_path.add_edges(3, 4)
	shortest_path.add_edges(7, 4)
	shortest_path.add_edges(7, 6)
	shortest_path.add_edges(6, 4)
	shortest_path.add_edges(6, 5)
	shortest_path.add_edges(4, 5)
	shortest_path.ShortestPath(0, 7)
}
