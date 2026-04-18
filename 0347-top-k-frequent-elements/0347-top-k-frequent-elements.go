func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)

    h := make(map[int]int)

    b := make([][]int, len(nums) + 1)

    for _, n := range nums {
        h[n]++
    }

    for n, c := range h {
        b[c] = append(b[c], n)
    }

    for i := len(b) - 1; i >= 0 && len(res) < k; i-- {
        if len(b[i]) > 0 {
            for _, n := range b[i] {
                res = append(res, n)
                if len(res) == k {
                    break
                }
            }
        }
    }
    
    return res
}