package main

import "fmt"

func ValidateChar(input string) bool {
	chars := []rune(input)
	allowedChars := map[rune]int{
		'<': 0, // 0 off 1 on
		'{': 0,
		'[': 0,
		'>': 0,
		'}': 0,
		']': 0,
	}

	if len(chars) == 0 && len(chars) > 4096{
		return false
	}

	if len(chars) == 1 && (chars[0] == '>' || chars[0] == '}' || chars[0] == ']'){
		return false
	}

	if chars[0] == '>' || chars[0] == '}' || chars[0] == ']' {
		if chars[1] == '<' || chars[1] == '[' || chars[1] == '{' {
			return false
		}
	}

	prevChars := []rune{}

	for _, char := range chars {
		_, exists := allowedChars[char]
		if !exists {
			return false
		}
		// set to on
		allowedChars[char] = 1
		if len(prevChars) > 0 {
			lastChart := prevChars[len(prevChars)-1]

			// if lastChart == '>' && char != '<'  {
			// 	return false
			// } else if lastChart == '}' && char != '{' {
			// 	return false
			// } else if lastChart == ']' && char != '[' {
			// 	return false
			// }

			if lastChart == '<' && (char == '}' || char == ']') {
				return false
			} else if lastChart == '{' && (char == '>' || char == ']') {
				return false
			} else if lastChart == '[' && (char == '>' || char == '}') {
				return false
			}
		}

		prevChars = append(prevChars, char)
	}
	return true
}

func main() {
	input := "]"
	result := ValidateChar(input)
	fmt.Println(result)
}
