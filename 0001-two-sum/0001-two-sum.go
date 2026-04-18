func twoSum(nums []int, target int) []int {
    h := make(map[int]int)

    for i, n := range nums {
        need := target - n

        if j, ok := h[need]; ok {
            return []int{i,j}
        }

        h[n] = i
    }
    return nil
}