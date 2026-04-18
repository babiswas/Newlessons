package arr

import (
	"container/list"
	"fmt"
)

type MotherVertex struct {
	vertex int
	adj    map[int]*list.List
}

func (m *MotherVertex) minit(vertex int) {
	m.vertex = vertex
	m.adj = make(map[int]*list.List)
	for i := 0; i < m.vertex; i++ {
		m.adj[i] = list.New()
	}
}

func (m *MotherVertex) add_edges(i, j int) {
	_, ok := m.adj[i]
	if !ok {
		m.adj[i] = list.New()
	}
	m.adj[i].PushBack(j)
}

func (m *MotherVertex) mother_vertex_util(u int, visited []bool) {
	visited[u] = true
	temp_list := m.adj[u]
	for el := temp_list.Front(); el != nil; el = el.Next() {
		index := el.Value.(int)
		if !visited[index] {
			m.mother_vertex_util(index, visited)
		}
	}
}

func (m *MotherVertex) mother_vertex() {
	m_vertex := -1
	visited := make([]bool, m.vertex)
	for i := 0; i < m.vertex; i++ {
		visited[i] = false
	}
	for i := range m.vertex {
		if !visited[i] {
			m.mother_vertex_util(i, visited)
			m_vertex = i
		}
	}
	for i := 0; i < m.vertex; i++ {
		visited[i] = false
	}
	m.mother_vertex_util(m_vertex, visited)
	for i := range m.vertex {
		if visited[i] == false {
			fmt.Println("No mother vertex in the graph")
			return
		}
	}
	fmt.Println("Mother vertex in the graph is:", m_vertex)
}

func Mother_Vertex() {
	fmt.Println("Mother vertex of the graph:")
	mv := MotherVertex{}
	mv.minit(4)
	fmt.Println("Constructing the graph:")
	mv.add_edges(0, 2)
	mv.add_edges(2, 0)
	mv.add_edges(0, 1)
	mv.add_edges(1, 2)
	mv.add_edges(2, 3)
	mv.add_edges(3, 3)
	fmt.Println("The mother vertex of the graph is:")
	mv.mother_vertex()
}
