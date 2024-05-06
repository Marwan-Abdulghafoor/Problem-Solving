package main

func main() {
	println(isValid("()[]{}"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if len(stack) != 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == '}' {
			if len(stack) != 0 && stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == ']' {
			if len(stack) != 0 && stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(s[i]))
		}
	}
	return len(stack) == 0
}
