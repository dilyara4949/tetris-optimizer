package main

import (
    "fmt"
)

func checkBlocks(blocks []string) error {
    if len(blocks)%4 != 0 {
        return fmt.Errorf("wrong size for block")
    }
    for _, line := range blocks {

        if len(line) != 4 {
            return fmt.Errorf("wrong size for line of block")
        }

        for _, symbol := range line {
            if symbol != '.' && symbol != '#' {
                return fmt.Errorf("wrong symbol inside block")
            }
        }

    }
    return nil
}

func checkTets(block []string) error {
    allAdj := 0 // cnt for adjacents if pieces of tetrominos
    cub := 0
    for i, line := range block {
        for j, symb := range line {
            if symb == '#' {
                cub++
                cntAdj := checkAdjacent(block, i, j)
                if cntAdj == 0 || cntAdj > 3 {
                    return fmt.Errorf("wrong tetromino")
                }
                allAdj += cntAdj
            }
        }
    }
    if (allAdj == 6  || allAdj == 8) && cub == 4 {
        return nil
    }
    return fmt.Errorf("wrong tetromino")
}

func checkAdjacent(block []string, i, j int) int {
    cntAdj := 0

    if i-1 >= 0 && block[i-1][j] == '#' { // up
        cntAdj++
    }
    if i+1 < 4 && block[i+1][j] == '#' { // down
        cntAdj++
    }
    if j-1 >= 0 && block[i][j-1] == '#' { // left
        cntAdj++
    }
    if j+1 < 4 && block[i][j+1] == '#' { // right
        cntAdj++
    }
    return cntAdj
}