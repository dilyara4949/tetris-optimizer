package main

import (
	"fmt"
	"os"
)

func placeTets(tets [][][]rune, Field [][]rune, size, tetCnt int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if checkPosition(tets[tetCnt], Field, i, j) {

				putToPosition(tets[tetCnt], Field, i, j, tetCnt+1)
				
				if tetCnt == len(tets)-1  {
					printAns(Field)
					
					os.Exit(0)
				} else {
					placeTets(tets, Field, size, tetCnt+1)
					deleteFromPosition(tets[tetCnt], Field, i, j, tetCnt)
				}
				

			}
		}
	}
	if tetCnt == 0 {
		size++

		Field = buildField(size)
		placeTets(tets, Field, size, tetCnt)
	}
	// x := 0
	// y := 0
	// n := len(tets[tetCnt])
	// m := len(tets[tetCnt][0])
	// fmt.Println(n, m)
	// fmt.Println(tets[tetCnt])
	// for i := x; i < x+n && i < len(Field); i++ {

	// 	for j := y; j < y+m && j < len(Field); j++ {

	// 		if (Field[i][j] == '.' && tets[tetCnt][i-x][j-y] == '#') {
	// 			// return false
	// 			fmt.Println(i, j, "   ", i-x, j-y)
	// 		}

	// 		// fmt.Println(i, j, "   ", i-x, j-y)

	// 	}
	// 	fmt.Println()
	// }
}

func putToPosition(tet [][]rune, Field [][]rune, x, y, tetCnt int) {
	n := len(tet)
	m := len(tet[0])
	for i := x; i < x+n && i < len(Field); i++ {

		for j := y; j < y+m && j < len(Field); j++ {

			if Field[i][j] == '.' && tet[i-x][j-y] == '#' {
				Field[i][j] = rune('A' + tetCnt-1)
			}

		}
	}

}

func deleteFromPosition(tet [][]rune, Field [][]rune, x, y, tetCnt int) {
	n := len(tet)
	m := len(tet[0])
	for i := x; i < x+n && i < len(Field); i++ {

		for j := y; j < y+m && j < len(Field); j++ {

			if tet[i-x][j-y] == '#' {
				Field[i][j] = '.'
			}

		}
	}

}

func checkPosition(tet [][]rune, Field [][]rune, x, y int) bool {
	n := len(tet)
	m := len(tet[0])
	cntSymbol := 0
	for i := x; i < x+n && i < len(Field); i++ {

		for j := y; j < y+m && j < len(Field); j++ {

			if i >= len(Field) || j >= len(Field)  {
				return false
			}
			if Field[i][j] == '.' && tet[i-x][j-y] == '#' {
				cntSymbol++
			}
		}
	}
	if cntSymbol == 4 {
		return true
	}
	return false

}

func printAns(Field [][]rune) {
	for _, line := range Field {
		for _, c := range line {
			fmt.Print((string(c)))
		}
		fmt.Println()
	}
}