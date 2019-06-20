package main

import "fmt"

func LevenshteinDistance(str1, str2 string) int {
	rowSize := len(str1) + 1
	columnSize := len(str2) + 1
	d := make([][]int, rowSize)
	for i := 0; i < rowSize; i++ {
		c := make([]int, columnSize)
		d[i] = c
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
				if str1[i-1] == str2[j-1] {
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
}

func main() {
	fmt.Println(LevenshteinDistance("aaaaa", "aacacccc"))
}
