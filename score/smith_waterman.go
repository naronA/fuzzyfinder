package score

func SmithWaterman(str1, str2 string) {
	// s1 := []rune(str1)
	// s2 := []rune(str2)

	// gap = 2
	// match = 2
	// misMatch - 2

	// n := len(s1) + 1
	// m := len(s2) + 1
	// mat, cmat := initMatrix(n, m)
	// mat[0][0] = 0
	// cmat[0][0] = NONE

	// for i := 1; i < m; i++ {
	// 	for j := 1; j < n; j++ {
	// 		di := mat[i-1][j-1] + diagonal(s1[j-1], s2[i-1])
	// 		ho := mat[i][j-1] - 2
	// 		ve := mat[i-1][j] - 2
	// 		mat[i][j] = max(di, ho, ve)
	// 		cmat[i][j] = pointers(di, ho, ve)
	// 	}
	// }
	// if draw {
	// 	printIntMat(s1, s2, mat)
	// 	printPointer(s1, s2, cmat)
	// 	drawResult(s1, s2, cmat)
	// }
	// return mat[m-1][n-1]
}
