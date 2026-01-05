package algo

import "container/list"

func MaxSlidingWindow(arr []int, K int) []int {
	start_index := 0
	end_index := 0
	queue := list.New()
	temp_list := make([]int, len(arr))
	temp_index := 0
	for end_index < len(arr) {
		for queue.Len() > 0 && queue.Front().Value.(int) < arr[end_index] {
			queue.Remove(queue.Front())
		}
		queue.PushBack(arr[end_index])
		window := end_index - start_index + 1
		if window == K {
			temp_num := queue.Front().Value.(int)
			temp_list[temp_index] = temp_num
			temp_index += 1
			if arr[start_index] == temp_num {
				queue.Remove(queue.Front())
			}
			start_index += 1
		}
		end_index += 1
	}
	return temp_list
}
