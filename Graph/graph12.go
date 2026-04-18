package arr

import (
	"fmt"
	"math"
)

func Dkstra(M [][]int, src int, vertex int) {
	fmt.Println("Dkstra algorithm:")
	parent := make([]int, vertex)
	for i := 0; i < vertex; i++ {
		parent[i] = -1
	}
	visited := make([]bool, vertex)
	for j := 0; j < vertex; j++ {
		visited[j] = false
	}
	cost := make([]int, vertex)
	for i := 0; i < vertex; i++ {
		cost[i] = math.MaxInt
	}
	find_minmum_cost_element := func(visited []bool, cost []int, vertex int) int {
		minm := math.MaxInt
		index := -1
		for i := range vertex {
			if cost[i] < minm && !visited[i] {
				index = i
				minm = cost[i]
			}
		}
		fmt.Println("Index is:", index)
		return index
	}
	cost[src] = 0
	for i := range vertex {
		min_element := find_minmum_cost_element(visited, cost, vertex)
		for j := range vertex {
			if cost[min_element]+M[min_element][j] < cost[j] && M[min_element][j] != 0 && !visited[j] {
				cost[j] = cost[min_element] + M[min_element][j]
				parent[j] = min_element
			}
		}
		visited[min_element] = true
		fmt.Println("Pass:", i)
	}
	fmt.Println("Parent nodes:")
	fmt.Println(parent)
	fmt.Println("Cost:")
	fmt.Println(cost)
}

func DkstraAlgorithm(vertex int) {
	fmt.Println("Single source shortest path in dkstra:")
	M := make([][]int, vertex)
	for i := 0; i < vertex; i++ {
		M[i] = make([]int, vertex)
	}
	cost1 := []int{0, 4, 0, 0, 0, 0, 0, 8, 0}
	cost2 := []int{4, 0, 8, 0, 0, 0, 0, 11, 0}
	cost3 := []int{0, 8, 0, 7, 0, 4, 0, 0, 2}
	cost4 := []int{0, 0, 7, 0, 9, 14, 0, 0, 0}
	cost5 := []int{0, 0, 0, 9, 0, 10, 0, 0, 0}
	cost6 := []int{0, 0, 4, 14, 10, 0, 2, 0, 0}
	cost7 := []int{0, 0, 0, 0, 0, 2, 0, 1, 6}
	cost8 := []int{8, 11, 0, 0, 0, 0, 1, 0, 7}
	cost9 := []int{0, 0, 2, 0, 0, 0, 6, 7, 0}

	l := [][]int{cost1, cost2, cost3, cost4, cost5, cost6, cost7, cost8, cost9}

	for j := 0; j < vertex; j++ {
		copy(M[j], l[j])
	}
	fmt.Println(M)
	Dkstra(M, 0, vertex)
}
