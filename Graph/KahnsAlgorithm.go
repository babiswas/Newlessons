package arr

import (
	"container/list"
	"fmt"
)

type KahnsAlgo struct {
	vertex int
	adj    map[int]*list.List
}

func (k *KahnsAlgo) kahns_init(vertex int) {
	k.vertex = vertex
	k.adj = make(map[int]*list.List)
	for i := 0; i < k.vertex; i++ {
		k.adj[i] = list.New()
	}
}

func (k *KahnsAlgo) add_edges(i, j int) {
	_, ok := k.adj[i]
	if !ok {
		k.adj[i] = list.New()
	}
	k.adj[i].PushBack(j)
}

func (k *KahnsAlgo) kahnsalgorithm() {
	count := 0
	indegree := make([]int, k.vertex)
	for i := 0; i < k.vertex; i++ {
		temp_list := k.adj[i]
		for elm := temp_list.Front(); elm != nil; elm = elm.Next() {
			index := elm.Value.(int)
			indegree[index] += 1
		}
	}
	queue := list.New()
	for i := range k.vertex {
		if indegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		el := queue.Front()
		num := el.Value.(int)
		fmt.Println(num)
		count += 1
		queue.Remove(el)
		temp_list := k.adj[num]
		for tmp := temp_list.Front(); tmp != nil; tmp = tmp.Next() {
			nm := tmp.Value.(int)
			indegree[nm] -= 1
			if indegree[nm] == 0 {
				queue.PushBack(nm)
			}
		}
	}

	if count == k.vertex {
		fmt.Println("There is no cycle in the graph.")
	} else {
		fmt.Println("There is a cycle in the graph.")
	}

}

func Kahns_Algorithm() {
	fmt.Println("Kahns algorithm:")
	kahns_algo := KahnsAlgo{}
	kahns_algo.kahns_init(6)
	fmt.Println("Constructing the graph:")
	kahns_algo.add_edges(5, 0)
	kahns_algo.add_edges(4, 0)
	kahns_algo.add_edges(5, 2)
	kahns_algo.add_edges(2, 3)
	kahns_algo.add_edges(3, 1)
	kahns_algo.add_edges(4, 1)
	fmt.Println("Executing kahns algorithm:")
	kahns_algo.kahnsalgorithm()
}
