func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    maxarea := 0

    for i:=0 ; i < m ; i++ {
        for j:=0 ; j < n ; j++ {
            if grid[i][j] == 1 {
                area := dfs(grid , i , j)
                if maxarea < area {
                    maxarea = area
                }
            }
        }

    }

    return maxarea
}

func dfs(grid [][]int, i , j int) int {
    m := len(grid)
    n := len(grid[0])

    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
        return 0
    }


    grid[i][j] = 0
    area := 1

    area += dfs(grid , i-1 ,j)
    area += dfs(grid , i ,j+1)
    area += dfs(grid , i+1, j)
    area += dfs(grid , i, j-1)

    return area
}