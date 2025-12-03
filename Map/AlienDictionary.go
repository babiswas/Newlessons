package swindow

import (
	"container/list"
	"fmt"
	"strings"
)

func reverse(str []string) {
	start := 0
	end := len(str) - 1
	for start <= end {
		str[start], str[end] = str[end], str[start]
		start += 1
		end -= 1
	}
}

func is_cyclic_util(visited, rec_stack map[string]bool, str string, mp map[string]*list.List, newlist *[]string) bool {
	visited[str] = true
	rec_stack[str] = true
	lst := mp[str]
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		newvalue := elm.Value.(string)
		if visited[newvalue] && rec_stack[newvalue] {
			return true
		} else if !visited[newvalue] {
			fmt.Println(newvalue)
			if is_cyclic_util(visited, rec_stack, newvalue, mp, newlist) {
				return true
			}
		}
	}
	rec_stack[str] = false
	(*newlist) = append((*newlist), str)
	_ = newlist
	return false
}

func iscyclic(mp map[string]*list.List, newlist *[]string) bool {
	visited := make(map[string]bool)
	for key, _ := range mp {
		visited[key] = false
	}
	rec_stack := make(map[string]bool)
	for key, _ := range mp {
		rec_stack[key] = false
	}

	for key, _ := range mp {
		if !visited[key] {
			fmt.Println("Detecting cycle in the graph.")
			if is_cyclic_util(visited, rec_stack, key, mp, newlist) {
				return true
			}
		}
	}
	reverse(*newlist)
	return false
}

func AlienDict(words []string) string {
	minm := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	mp := make(map[string]*list.List)
	for _, val := range words {
		for _, ch := range strings.Split(val, "") {
			_, ok := mp[ch]
			if !ok {
				mp[ch] = list.New()
			}
		}
	}
	for i := 0; i < len(words)-1; i++ {
		wd1 := words[i]
		wd2 := words[i+1]
		minml := minm(len(wd1), len(wd2))
		if len(wd1) > len(wd2) && wd1[:minml] == wd2[:minml] {
			return ""
		}
		for j := range minml {
			if wd1[j] != wd2[j] {
				mp[string(wd1[j])].PushBack(string(wd2[j]))
				break
			}
		}
	}
	for key, val := range mp {
		fmt.Print(key, ":")
		for elm := val.Front(); elm != nil; elm = elm.Next() {
			fmt.Print(" ", elm.Value.(string), ",")
		}
		fmt.Println("")
	}
	newlist := make([]string, 0)
	if iscyclic(mp, &newlist) {
		fmt.Println(newlist)
		return ""
	}
	fmt.Println(newlist)
	return strings.Join(newlist, "")
}
