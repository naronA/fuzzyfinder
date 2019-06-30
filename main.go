package main

import (
	"fmt"
	"math"
)

func PrintRuneMat(mat [][]rune) {
	for _, col := range mat {
		for _, val := range col {
			fmt.Printf("%c\t", val)
		}
		fmt.Println()
		// fmt.Println(col)
	}
}

func PrintIntMat(mat [][]int) {
	for _, col := range mat {
		for _, val := range col {
			fmt.Printf("%v\t", val)
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

func diagonal(n1, n2 rune) int {
	if n1 == n2 {
		return 2
	}
	return -1
}

func pointers(di, ho, ve int) rune {
	pointer := max(di, ho, ve)
	if di == pointer {
		return 'D'
	} else if ho == pointer {
		return 'H'
	}
	return 'V'
}

func max(x ...int) int {
	max := math.MinInt64
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func NeedlemanWunsch(s1, s2 []rune) int {
	n := len(s1) + 1
	m := len(s2) + 1
	mat := make([][]int, m)
	cmat := make([][]rune, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
		cmat[i] = make([]rune, n)
	}

	for i := 0; i < m; i++ {
		mat[i][0] = i * -2
		cmat[i][0] = 'V'
	}

	for j := 0; j < n; j++ {
		mat[0][j] = j * -2
		cmat[0][j] = 'H'
	}

	mat[0][0] = 0
	cmat[0][0] = 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			di := mat[i-1][j-1] + diagonal(s1[j-1], s2[i-1])
			ho := mat[i][j-1] - 2
			ve := mat[i-1][j] - 2
			mat[i][j] = max(di, ho, ve)
			cmat[i][j] = pointers(di, ho, ve)
		}
	}
	PrintIntMat(mat)
	fmt.Println()
	PrintRuneMat(cmat)
	return mat[m-1][n-1]
}

func main() {
	score := NeedlemanWunsch([]rune("ATACCC"), []rune("ACBCC"))
	fmt.Println(score)
}
