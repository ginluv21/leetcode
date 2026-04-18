func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    h := make(map[rune]int)

    for _, ch := range s {
        h[ch]++
    }

    for _, ch := range t {
        h[ch]--

        if h[ch] < 0 {
            return false
        }
    }

    return true
}