func longestCommonSubsequence(text1 string, text2 string) int {
    n := len(text1)
    m := len(text2)
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, m + 1)
    }
    for i, c1 := range text1 {
        for j, c2 := range text2 {
            if c1 == c2 {
                dp[i+1][j+1] = dp[i][j] + 1
            } else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[n][m]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
