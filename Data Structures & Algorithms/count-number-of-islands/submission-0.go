func numIslands(grid [][]byte) int {

    visitedElement := map[string]bool{}

    result := 0

    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j] == '0' || visitedElement[matrixKey(i,j)] {
                continue
            }
    
            result += 1
            find(grid, visitedElement, i, j)

        }
    }

    return result
}

func matrixKey(row, column int) string {
    return fmt.Sprintf("%d:%d", row, column)
}

func find(grid [][]byte, visitedElement map[string]bool, pointerRow, pointerColumn int) {


    if pointerRow < 0 || pointerRow >= len(grid) || pointerColumn < 0 || pointerColumn >= len(grid[0]) {
        return
    }
    if grid[pointerRow][pointerColumn] == '0' {
        return
    }

    key := matrixKey(pointerRow, pointerColumn)
    if visitedElement[key] {
        return
    }
    visitedElement[key] = true
    fmt.Println(key)

    find(grid, visitedElement, pointerRow, pointerColumn-1)
    

    find(grid, visitedElement, pointerRow-1, pointerColumn)
    

    find(grid, visitedElement, pointerRow+1, pointerColumn)
    

    find(grid, visitedElement, pointerRow, pointerColumn+1)
} 