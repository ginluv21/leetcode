func twoSum(numbers []int, target int) []int {
    h:= make(map[int]int)
	
	for i, n := range numbers{
		need := target - n;

		if j, ok := h[need]; ok {
			return []int{j + 1, i + 1}
		}

		h[n] = i
	}
	return nil
}