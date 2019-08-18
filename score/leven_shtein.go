package score

func LevenshteinDistance(str1, str2 string) int {
	s1 := []rune(str1)
	s2 := []rune(str2)
	rowSize := len(s1) + 1
	columnSize := len(s2) + 1
	d := make([][]int, rowSize)
	for i := 0; i < rowSize; i++ {
		d[i] = make([]int, columnSize)
	}

	for i := 0; i < rowSize; i++ {
		d[i][0] = i
	}

	for j := 0; j < columnSize; j++ {
		d[0][j] = j
	}

	for i := 1; i < rowSize; i++ {
		for j := 1; j < columnSize; j++ {
			con := func() int {
				if s1[i-1] == s2[j-1] {
					return 0
				}
				return 1
			}()
			x1 := d[i-1][j] + 1
			x2 := d[i][j-1] + 1
			x3 := d[i-1][j-1] + con
			switch {
			case x1 <= x2 && x1 <= x3:
				d[i][j] = x1
			case x2 <= x1 && x2 <= x3:
				d[i][j] = x2
			case x3 <= x1 && x3 <= x2:
				d[i][j] = x3
			}
		}
	}
	return d[rowSize-1][columnSize-1]
	// , []int{}
}
