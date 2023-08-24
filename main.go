package main

import (
	"fmt"
)

func isBalanced(str string) bool {
	stack := []rune{}
	steplesMap := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, char := range str {
		// Если это открывающая скобка, добавляем ее в стек
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else if char == '}' || char == ']' || char == ')' {
			// Если это закрывающая скобка, проверяем, есть ли соответствующая открывающая скобка на вершине стека
			if len(stack) == 0 || stack[len(stack)-1] != steplesMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	input := "{[()()]}"
	if isBalanced(input) {
		fmt.Println("Строка сбалансированная")
	} else {
		fmt.Println("Строка не сбалансированная")
	}
}
