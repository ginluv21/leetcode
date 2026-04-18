func reverse(n []int, l, r int) {

	for ; l < r; l, r = l+1, r-1 {
		n[l], n[r] = n[r], n[l]
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}

    k = k % l

	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)

}