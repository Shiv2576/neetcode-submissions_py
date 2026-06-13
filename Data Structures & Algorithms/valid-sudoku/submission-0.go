func isValidSudoku(board [][]byte) bool {
    m := len(board)

    for i:=0 ; i < m ; i++ {
        seen := make(map[byte]bool)
        for j := 0 ; j< m ; j++ {

            val  := board[i][j]

            if board[i][j] == '.' {
                continue
            }


            if seen[val] {
                return false
            } 
               
            seen[val] = true
        }
    }

    for i:=0 ; i < m ; i++ {
        seen := make(map[byte]bool)
        for j := 0 ; j< m ; j++ {
            val := board[j][i]
            if board[j][i] == '.' {
                continue
            }

            if seen[val] {
                return false
            } 
            seen[val] = true
        }
    }


    for square:=0 ; square < 9 ; square++ {
        seen := make(map[byte]bool)

        squarerow := (square / 3) * 3
        squarecol := (square % 3) * 3

        for i:= 0 ; i < 3 ; i++ {
            for j := 0 ; j < 3 ; j++ {
                r := squarerow + i
                c := squarecol + j

                val := board[r][c]
                if val == '.' {
                    continue
                }

                if seen[val] {
                    return false
                }

                seen[val] = true
            }
        }
    }


    return true




}
