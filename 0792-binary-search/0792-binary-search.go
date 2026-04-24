func search(nums []int, target int) int {
    r, l := 0, len(nums) - 1

    for r <= l {
        mid := l + (r - l)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            r = mid + 1
        } else {
            l = mid - 1
        }
    }

    return -1
}