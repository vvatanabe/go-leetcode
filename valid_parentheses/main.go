package main

func isValid(s string) bool {
	temp := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	var stack []string
	for i := 0; i < len(s); i++ {
		if _, ok := temp[s[i:i+1]]; ok {
			stack = append(stack, s[i:i+1])
			continue
		}
		if len(stack) == 0 || temp[stack[len(stack)-1]] != s[i:i+1] {
			return false
		}
		stack = stack[0 : len(stack)-1]
	}
	return len(stack) == 0

}

func isValid2(s string) bool {
	temp := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	var stack []string
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			if _, ok := temp[s[i:i+1]]; !ok {
				return false
			}
			stack = append(stack, s[i:i+1])
			continue
		}
		if _, ok := temp[s[i:i+1]]; ok {
			stack = append(stack, s[i:i+1])
			continue
		}
		if temp[stack[len(stack)-1]] != s[i:i+1] {
			return false
		}
		stack = stack[0 : len(stack)-1]
	}
	return len(stack) == 0
}
