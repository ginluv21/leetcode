func containsDuplicate(nums []int) bool {
    h := make(map[int]bool)

    for _, n := range nums {
        if h[n] {
            return true
        }
        h[n] = true
    }
    return false
}