func isValidSudoku(board [][]byte) bool {
	mapValid := map[byte]bool{}
	// check row
	row := 0
	column := 0
	for {
		if row == len(board) {
			mapValid = map[byte]bool{}
			break
		}
		if column == len(board[0]) {
			column = 0
			row+=1
			mapValid = map[byte]bool{}
			continue
		}
		if board[row][column] == '.' {
			column += 1
			continue
		}
		if _, ok := mapValid[board[row][column]]; ok {
			return false
		}
		mapValid[board[row][column]] = true
		column+=1
	}

	// check column
	row = 0
	column = 0
	for {
		if column == len(board[0]) {
			mapValid = map[byte]bool{}
			break
		}
		if row == len(board) {
			column+=1
			row = 0
			mapValid = map[byte]bool{}
			continue
		}
		if board[row][column] == '.' {
			row += 1
			continue
		}
		if _, ok := mapValid[board[row][column]]; ok {
			return false
		}
		mapValid[board[row][column]] = true
		row+=1
	}

	// check subsudoku
	
	// i = row - (row%3)
	// j := column - (column%3)
	// maxRow := i+3
	// maxColumn := j+3
	row = 0
	column = 0

	for {
		mapValid = map[byte]bool{}

		if row == len(board) {
			break
		}

		maxRow := row+3
		maxColumn := column+3
		for i:=row;i<maxRow;i++{
			for j:=column;j<maxColumn;j++ {
				if board[i][j] == '.' {
					continue
				}
				if _, ok := mapValid[board[i][j]]; ok {
					return false
				}
				mapValid[board[i][j]] = true
			}
		}

		column+=3
		if column < len(board[0]) {
			continue
		} else {
			row = row + 3
			column = 0
		}
	}



	return true
}
