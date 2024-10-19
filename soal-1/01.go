package main

import "fmt"

func CountPairString(n int, str string) string {
	stringToSlice := []string{}
	var tmp string
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			tmp += string(str[i])
		} else {
			stringToSlice = append(stringToSlice, tmp)
			tmp = ""
		}
	}

	if tmp != "" {
		stringToSlice = append(stringToSlice, tmp)
	}

	pairs := make(map[string]struct{
		value int
		index []int
	})
	for index, val := range stringToSlice {
		pair, isExist := pairs[val]
		if isExist {
			pair.index = append(pair.index, index)
			pair.value += 1
			pairs[val] = pair 
		} else {
			pairs[val] = struct {
				value int
				index []int
			}{
				value: 1,         
				index: []int{index},  
			}
		}
	}

	result := ""
	for _, pair := range pairs {
		if pair.value > 1 {
			for _, v := range pair.index{
				result += fmt.Sprintf("%d ", v + 1)
			}
		}
	}

	if result == "" {
		return "false"
	}

	return result
}


