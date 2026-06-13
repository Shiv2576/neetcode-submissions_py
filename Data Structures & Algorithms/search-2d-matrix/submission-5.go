func searchMatrix(matrix [][]int, target int) bool {

    m := len(matrix)
    n := len(matrix[0])

    r := 0



    for i := 0 ; i < m ; i++ {
        if matrix[i][n-1] == target {
            return true
        } else if matrix[i][n-1] > target {
            r = i
            break
        }
    }

    for i := 0 ; i < n ; i++ {
        if matrix[r][i] == target {
            return true
        }
    }

    return false

}
