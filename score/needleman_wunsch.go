package score

import (
	"math"

	"github.com/naronA/fuzzyfinder/config"
)

const (
	GAP      = 2
	MATCH    = 2
	MISMATCH = -1
)

func initNeedlemanWunsch(n, m int) ([][]int, [][]int) {
	mat := make([][]int, m)
	cmat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
		cmat[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		mat[i][0] = i * -2
		cmat[i][0] = V
	}

	for j := 0; j < n; j++ {
		mat[0][j] = j * -2
		cmat[0][j] = H
	}
	return mat, cmat
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

func diagonal(n1, n2 rune) int {
	if n1 == n2 {
		return MATCH
	}
	return MISMATCH
}

func pointers(di, ho, ve int) int {
	pointer := max(di, ho, ve, 0)
	switch {
	case pointer == 0:
		return DHV
	case di == pointer && ho == pointer && ve == pointer:
		return DHV
	case di == pointer && ho == pointer:
		return DH
	case di == pointer && ve == pointer:
		return DV
	case ho == pointer && ve == pointer:
		return HV
	case ve == pointer:
		return V
	case ho == pointer:
		return H
	case di == pointer:
		return D
	}
	return V
}

func NeedlemanWunsch(str1, str2 string) int {
	s1 := []rune(str1)
	s2 := []rune(str2)
	n := len(s1) + 1
	m := len(s2) + 1
	mat, cmat := initNeedlemanWunsch(n, m)
	mat[0][0] = 0
	cmat[0][0] = NONE
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			di := mat[i-1][j-1] + diagonal(s1[j-1], s2[i-1])
			ho := mat[i][j-1] - GAP
			ve := mat[i-1][j] - GAP
			mat[i][j] = max(di, ho, ve)
			cmat[i][j] = pointers(di, ho, ve)
		}
	}
	if config.DEBUG {
		printIntMat(s1, s2, mat)
		printPointer(s1, s2, cmat)
		drawResult(s1, s2, cmat)
	}

	return mat[m-1][n-1]
	//, matched(s1, s2, cmat)
}
