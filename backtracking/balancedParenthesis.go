package arr

import (
	"container/list"
)

func BalancedParenthesis(parenthesis string) bool {
	is_exist := func(arr []string, ch string) bool {
		for _, val := range arr {
			if val == ch {
				return true
			}
		}
		return false
	}
	openParenthesis := []string{"{", "(", "["}
	closedParenthesis := []string{"}", ")", "]"}
	pmap := make(map[string]string)
	pmap = map[string]string{"{": "}", "(": ")", "[": "]"}
	stack := list.New()
	for _, val := range parenthesis {
		temp := string(val)
		if stack.Len() == 0 {
			if is_exist(openParenthesis, temp) {
				stack.PushBack(temp)
			} else {
				return false
			}
		} else {
			elm := stack.Back()
			topval := elm.Value.(string)
			exist1 := is_exist(openParenthesis, temp)
			exist2 := is_exist(closedParenthesis, temp)
			if exist1 {
				stack.PushBack(temp)
			} else if exist2 {
				if pmap[topval] == temp {
					stack.Remove(elm)
				} else {
					return false
				}
			}

		}
	}
	if stack.Len() == 0 {
		return true
	} else {
		return false
	}

}
