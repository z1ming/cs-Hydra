func lengthOfLongestSubstring(s string) int {
    n := len(s)
    set := map[byte]int{}
    right, ans := -1, 0
    for i := 0; i < n; i++ {
        if i != 0 {
            delete(set, s[i - 1])
        }
        for right + 1 < n && set[s[right + 1]] == 0 {
            set[s[right + 1]]++
            right++
        }
        ans = max(ans, right - i + 1)
    }
    return ans
}

func max(a, b int) int {
    if (a < b) {
        return b
    }
    return a
}
