package arr

import (
	"container/list"
	"fmt"
)

type Position struct {
	xpos int
	ypos int
}

var paths = list.New()

func is_valid_path(x, y int, visited [][]bool, N int, M [][]int) bool {
	if (x < N && y < N) && (x >= 0 && y >= 0) && M[x][y] != 0 && !visited[x][y] {
		return true
	}
	return false
}

func find_path(x, y int, visited [][]bool, M [][]int, N int, targetx, targety int, temp []Position) {
	if x == targetx && y == targety {
		paths.PushBack(temp)
		return
	}
	choices := []Position{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}
	for _, pos := range choices {
		if is_valid_path(pos.xpos, pos.ypos, visited, N, M) {
			visited[pos.xpos][pos.ypos] = true
			temp = append(temp, Position{pos.xpos, pos.ypos})
			find_path(pos.xpos, pos.ypos, visited, M, N, targetx, targety, temp)
			visited[pos.xpos][pos.ypos] = false
			temp = temp[:len(temp)-1]
		}
	}
}

func FindRatsPathMaze() {
	fmt.Println("Find the path of the rat in maze:")
	visited := make([][]bool, 4)
	for i := 0; i < 4; i++ {
		visited[i] = make([]bool, 4)
		for j := 0; j < 4; j++ {
			visited[i][j] = false
		}
	}
	M1 := []int{1, 0, 0, 0}
	M2 := []int{1, 1, 0, 1}
	M3 := []int{1, 1, 0, 0}
	M4 := []int{0, 1, 1, 1}
	M := [][]int{M1, M2, M3, M4}
	temp := make([]Position, 0)
	temp = append(temp, Position{0, 0})
	find_path(0, 0, visited, M, 4, 3, 3, temp)
	for el := paths.Front(); el != nil; el = el.Next() {
		pth := el.Value.([]Position)
		fmt.Println(pth)
	}
}
