func removeDuplicates(nums []int) int {
	res := 1

	s, f := 0, 0

	for f < len(nums){
		if nums[f] != nums[s] {
			nums[s + 1] = nums[f]
			s++
			res++
		} else {
			f++
		}
	}

	return res
    
}