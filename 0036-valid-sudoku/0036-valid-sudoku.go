func isValidSudoku(board [][]byte) bool {
	var r, c, b [9][9]bool

	for i:= 0; i < 9; i++ {
		for j:= 0; j < 9; j++{
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1'
			box := (i /3)*3 + j/3

			if r[i][num] || c[j][num] || b[box][num] {
				return false
			}

			r[i][num] = true 
			c[j][num] = true
			b[box][num] = true
		}
	}

	return true
}