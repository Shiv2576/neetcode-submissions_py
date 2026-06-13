func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])

    count := 0

    for i := 0 ; i < m ; i++ {
        for j := 0 ; j < n ; j++ {
            if grid[i][j] == '1' {
                count ++
                dfs(grid, i ,j)
            }
        }
    }

    return count
    
}


func dfs(grid [][]byte, i , j int) {
    m := len(grid)
    n := len(grid[0])

    if i < 0 || i >= m || j < 0 || j >= n {
        return
    }

    if grid[i][j] == '0' {
        return
    }


    grid[i][j] = '0'


    dfs(grid,i-1,j)
    dfs(grid,i,j+1)
    dfs(grid,i+1,j)
    dfs(grid,i,j-1)
}