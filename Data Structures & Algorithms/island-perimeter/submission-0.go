func islandPerimeter(grid [][]int) int {

    perimeter := 0

    rows := []int{-1, 1, 0, 0}
    cols := []int{0, 0, -1, 1}

    for i := 0; i < len(grid); i++ {

        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            }

            for k := 0; k < 4; k++ {

                newRow := i + rows[k]
                newCol := j + cols[k]

                if newRow < 0 ||
                   newRow >= len(grid) ||
                   newCol < 0 ||
                   newCol >= len(grid[0]) {

                    perimeter++
                    continue
                }

                if grid[newRow][newCol] == 0 {
                    perimeter++
                }
            }
        }
    }

    return perimeter
}