package chapter_5

func SudokuChecker(board [][]int) bool {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			curr := board[x][y]
			if curr == 0 {
				continue
			}

			// Check column
			for i := 0; i < len(board); i++ {
				if board[i][y] == curr && i != x {
					return false
				}
			}

			// Check row
			for j := 0; j < len(board); j++ {
				if board[x][j] == curr && j != y {
					return false
				}
			}

			// Check 3x3 Subarray
			xOffset := x / 3
			yOffset := y / 3
			for i := xOffset * 3; i <= xOffset*3+2; i++ {
				for j := yOffset * 3; j <= yOffset*3+2; j++ {
					if board[i][j] == curr && i != x && j != y {
						return false
					}
				}
			}
		}
	}
	return true
}
