func isValid(s string) bool {
	stack := []rune{}

	hash := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {

		if r, ok := hash[ch]; ok {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack) -1]
			stack = stack[:len(stack) -1]

			if top != r{
				return false
			}
		} else {
			stack = append(stack, ch)
		}
	} 

	return len(stack) == 0
}