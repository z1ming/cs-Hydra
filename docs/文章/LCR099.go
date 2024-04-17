func minPathSum(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    f := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, m)
    }
    f[0][0] = grid[0][0]
    for i := 1; i < n; i++ {
        f[i][0] = f[i-1][0] + grid[i][0]
    }
    for i := 1; i < m; i++ {
        f[0][i] = f[0][i-1] + grid[0][i]
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
        }
    }
    return f[n-1][m-1]
}

func min(a, b int) int {
    if (a < b) {return a}
    return b
}
