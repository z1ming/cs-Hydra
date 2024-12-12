func maxAreaOfIsland(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    ans := 0
    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 &&!visited[i][j] {
                ans = max(ans, getArea(grid, visited, i, j))
            }
        }
    }
    return ans
}

func getArea(grid [][]int, visited [][]bool, i, j int) int {
    area := 1
    visited[i][j] = true
    dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    for _, dir := range dirs {
        r := i + dir[0]
        c := j + dir[1]
        if (r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 && !visited[r][c]) {
            area += getArea(grid, visited, r, c)
        }
    }
    return area
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
