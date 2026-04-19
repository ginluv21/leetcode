func removeElement(nums []int, val int) int {
    s, f := 0, 0 

	for ; f < len(nums); f++ {
		if nums[s] == val && nums[f] != val {
			nums[s], nums[f] = nums[f], nums[s]
			s++
		}

		if nums[s] != val {
			s++
		}
	}

	return s
}