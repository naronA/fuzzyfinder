package main

import "fmt"

func PrintMat(mat [][]int) {
	for _, col := range mat {
		for _, val := range col {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
		// fmt.Println(col)
	}
}

func initNedlemanWunsch(row, column int) [][]int {
	mat := make([][]int, row)
	for i := 0; i < row; i++ {
		mat[i] = make([]int, column)
	}

	for i := 0; i < row; i++ {
		mat[i][0] = i * -2
	}

	for j := 0; j < column; j++ {
		mat[0][j] = j * -2
	}
	return mat
}

func NeedlemanWunsch(s1, s2 []rune) int {
	row := len(s1) + 1
	column := len(s2) + 1
	mat := initNedlemanWunsch(row, column)
	PrintMat(mat)
	return 0
}

func main() {
	NeedlemanWunsch([]rune("aaaa"), []rune("bbbb"))
}
