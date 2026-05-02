func numIslands(grid [][]byte) int {

    result := 0

    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j] == '0' {
                continue
            }
    
            result += 1
            find(grid,  i, j)

        }
    }

    return result
}

func matrixKey(row, column int) string {
    return fmt.Sprintf("%d:%d", row, column)
}

func find(grid [][]byte, pointerRow, pointerColumn int) {

    if pointerRow < 0 || pointerRow >= len(grid) || pointerColumn < 0 || pointerColumn >= len(grid[0]) {
        return
    }
    if grid[pointerRow][pointerColumn] == '0' {
        return
    }

    grid[pointerRow][pointerColumn] = '0'

    find(grid, pointerRow, pointerColumn-1)
    

    find(grid, pointerRow-1, pointerColumn)
    

    find(grid, pointerRow+1, pointerColumn)
    

    find(grid, pointerRow, pointerColumn+1)
} 