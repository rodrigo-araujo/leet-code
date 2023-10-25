package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "()"
	result := isValid(s)
	fmt.Println("Result: ", result)

	s = "()[]{}"
	result = isValid(s)
	fmt.Println("Result: ", result)

	s = "(]"
	result = isValid(s)
	fmt.Println("Result: ", result)

}

func isValid(s string) bool {

	var stack Stack

	for x := 0; x < len(s); x++ {
		currentValue := string(s[x])
		if x == 0 || strings.Contains("{[(", currentValue) {
			stack.Push(currentValue)
			continue
		}

		prev, _ := stack.Pop()
		switch currentValue {
		case ")":
			if prev == "(" {
				continue
			}
			return false
		case "]":
			if prev == "[" {
				continue
			}
			return false
		case "}":
			if prev == "{" {
				continue
			}
			return false
		}
	}

	return stack.IsEmpty()
}

type Stack []string

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		value := (*s)[index]
		*s = (*s)[:index]
		return value, true
	}
}
