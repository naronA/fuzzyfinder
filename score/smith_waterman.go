package score

import (
	"github.com/naronA/fuzzyfinder/config"
	"strings"
)

func initSmithWaterman(n, m int) ([][]int, [][]int) {
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

func SmithWaterman(str1, str2 string) int {
	s1 := []rune(str1)
	s2 := []rune(str2)
	n := len(s1) + 1
	m := len(s2) + 1
	mat, cmat := initSmithWaterman(n, m)
	mat[0][0] = 0
	cmat[0][0] = NONE
	score := 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			di := mat[i-1][j-1] + diagonal(s1[j-1], s2[i-1])
			ho := mat[i][j-1] - GAP
			ve := mat[i-1][j] - GAP
			mat[i][j] = max(ho, ve, di, 0)
			score = max(score, mat[i][j])
			// cmat[i][j] = pointers(di, ho, ve)
		}
	}
	if config.DEBUG {
		printIntMat(s1, s2, mat)
		printPointer(s1, s2, cmat)
		drawResult(s1, s2, cmat)
	}
	if strings.Contains(str1, str2) {
		score += MATCH
	}

	return score
	// , matched(s1, s2, cmat)
}
