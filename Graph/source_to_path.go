package arr

import (
	"container/list"
	"fmt"
)

var all_path = make([][]int, 0)

type AllPath struct {
	vertex int
	adj    map[int]*list.List
}

func (a *AllPath) all_path_init(vertex int) {
	a.vertex = vertex
	a.adj = make(map[int]*list.List)
	for i := 0; i < a.vertex; i++ {
		a.adj[i] = list.New()
	}
}

func (a *AllPath) add_edges(i, j int) {
	_, ok := a.adj[i]
	if !ok {
		a.adj[i] = list.New()
	}
	a.adj[i].PushBack(j)
}

func (a *AllPath) find_path(src, target int) {
	path := make([]int, 0)
	visited := make([]bool, a.vertex)
	for i := 0; i < a.vertex; i++ {
		visited[i] = false
	}
	a.find_path_util(src, target, visited, path)
	fmt.Println(all_path)
}

func (a *AllPath) find_path_util(src, target int, visited []bool, path []int) {
	path = append(path, src)
	visited[src] = true
	if src == target {
		temp_path := make([]int, len(path))
		copy(temp_path, path)
		all_path = append(all_path, temp_path)
	} else {
		temp_list := a.adj[src]
		for el := temp_list.Front(); el != nil; el = el.Next() {
			index := el.Value.(int)
			if !visited[index] {
				a.find_path_util(index, target, visited, path)
			}

		}
	}
	visited[src] = false
	path = path[:len(path)-1]
}

func DisplayAllPath() {
	fmt.Println("Displaying all path from source to destination:")
	allPath := AllPath{}
	fmt.Println("Find all path from source to target:")
	allPath.all_path_init(4)
	allPath.add_edges(0, 2)
	allPath.add_edges(2, 0)
	allPath.add_edges(0, 1)
	allPath.add_edges(0, 3)
	allPath.add_edges(1, 2)
	allPath.add_edges(1, 3)
	allPath.add_edges(2, 3)
	allPath.add_edges(3, 3)
	fmt.Println("Finding paths:")
	allPath.find_path(2, 3)
}
