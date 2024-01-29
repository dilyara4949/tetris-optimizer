package main

import "math"

func getMinFiledSize(tetrominoes [][][]rune) int {
    n := math.Sqrt(float64(len(tetrominoes)*4))
    return int(math.Ceil(n))
}

func buildField(n int) [][]rune {
    field := [][]rune{}
    for i := 0; i < n; i++ {
        field = append(field, nil)
        for j := 0; j < n; j++ {
            field[i] = append(field[i], '.')
        }
    }
    return field
}