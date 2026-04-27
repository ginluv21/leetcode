func longestConsecutive(nums []int) int {
	h := make(map[int]bool)
	max := 0

	for _, n := range nums {
		h[n] = true
	}

	for n := range h {
		if h[n-1] {
			continue
		}

		l := 1
		c := n

		for h[c+1] {
			l++
			c++
		}

		if l > max {
			max = l
		}
	}
	return max
}