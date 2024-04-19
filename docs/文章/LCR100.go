func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    f := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, len(triangle[i]))
        if i == 0 {
            f[i][0] = triangle[0][0]
        } else {
            m := len(triangle[i])
            for j := 0; j < m; j++ {
                if j == 0 {
                    f[i][j] = f[i - 1][j] + triangle[i][j];
                } else if j == m - 1 {
                    f[i][j] = f[i-1][j-1] + triangle[i][j];
                } else {
                    f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j];
                }
            }
        }
    }
    ans := 10000
    for i := 0; i < len(f[n-1]); i++ {
        ans = min(ans, f[n-1][i])
        // fmt.Println(ans)
    }
    return ans
}

func min(a, b int) int {
    if (a < b) {
        return a
    }
    return b
}
