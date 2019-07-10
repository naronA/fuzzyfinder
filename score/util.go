package score

import (
	"fmt"
)

const (
	DHV = iota
	DH
	DV
	HV
	D
	H
	V
	NONE
)

func printPointer(s1, s2 []rune, mat [][]int) {
	fmt.Printf("%2s  |", " ")
	fmt.Printf("%5s", " ")
	for _, r := range s1 {
		fmt.Printf("%4c ", r)
	}
	fmt.Println()

	fmt.Print("----+")
	for i := 0; i < len(s1)*5+5; i++ {
		fmt.Print("-")
	}

	fmt.Println()
	for i, col := range mat {
		if i == 0 {
			fmt.Printf("%2s  |", " ")
		} else {
			fmt.Printf("%2c  |", s2[i-1])
		}

		for _, val := range col {
			var token string
			switch val {
			case 0:
				token = "DHV"
			case 1:
				token = "DH"
			case 2:
				token = "DV"
			case 3:
				token = "HV"
			case 4:
				token = "D"
			case 5:
				token = "H"
			case 6:
				token = "V"
			default:
				token = "0"
			}
			fmt.Printf("%4s ", token)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printIntMat(s1, s2 []rune, mat [][]int) {
	fmt.Printf("%2s  |", " ")
	fmt.Printf("%5s", " ")
	for _, r := range s1 {
		fmt.Printf("%4c ", r)
	}
	fmt.Println()

	fmt.Print("----+")
	for i := 0; i < len(s1)*5+5; i++ {
		fmt.Print("-")
	}

	fmt.Println()
	for i, col := range mat {
		if i == 0 {
			fmt.Printf("%2s  |", " ")
		} else {
			fmt.Printf("%2c  |", s2[i-1])
		}

		for _, val := range col {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func drawResult(s1, s2 []rune, cmat [][]int) {
	n := len(s1) + 1
	m := len(s2) + 1
	x := m - 1
	y := n - 1

	if n > m {
		long := make([]rune, n)
		mid := make([]rune, n)
		short := make([]rune, n)
		pos := n - 2
		shortPos := m - 2
		for x != 0 || y != 0 {
			if pos < 0 {
				break
			}

			switch cmat[x][y] {
			case D:
				if s1[pos] == s2[shortPos] {
					x--
					y--
					long[pos] = s1[pos]
					mid[pos] = '|'
					short[pos] = s2[shortPos]
					shortPos--
				} else {
					x--
					y--
					long[pos] = s1[pos]
					mid[pos] = ' '
					short[pos] = '-'
				}
			case DHV, DH, HV, H:
				y--
				long[pos] = s1[pos]
				mid[pos] = ' '
				short[pos] = '-'
			case DV, V:
				x--
				long[pos] = s1[pos]
				mid[pos] = ' '
				short[pos] = '-'
			}
			pos--
		}
		fmt.Printf("%s\n%s\n%s\n", string(long), string(mid), string(short))
		fmt.Println()
		return
	}
	long := make([]rune, m)
	mid := make([]rune, m)
	short := make([]rune, m)
	pos := m - 2
	shortPos := n - 2
	for x != 0 || y != 0 {
		if pos < 0 {
			break
		}
		switch cmat[x][y] {
		case D:
			if s2[pos] == s1[shortPos] {
				x--
				y--
				long[pos] = s2[pos]
				mid[pos] = '|'
				short[pos] = s1[shortPos]
				shortPos--
			} else {
				x--
				y--
				long[pos] = s2[pos]
				mid[pos] = ' '
				short[pos] = '-'
			}
		case DH, H:
			y--
			long[pos] = s2[pos]
			mid[pos] = ' '
			short[pos] = '-'
		case DHV, HV, DV, V:
			x--
			long[pos] = s2[pos]
			mid[pos] = ' '
			short[pos] = '-'
		}
		pos--
	}

	fmt.Printf("%s\n%s\n%s\n", string(long), string(mid), string(short))
	fmt.Println()
}

func matched(s1, s2 []rune, cmat [][]int) []int {
	n := len(s1) + 1
	m := len(s2) + 1
	x := m - 1
	y := n - 1
	matchedPos := []int{}
	if n > m {
		pos := n - 2
		shortPos := m - 2
		for x != 0 || y != 0 {
			if pos < 0 {
				break
			}

			switch cmat[x][y] {
			case D:
				x--
				y--
				if s1[pos] == s2[shortPos] {
					matchedPos = append(matchedPos, pos)
					shortPos--
				}
			case DHV, DH, HV, H:
				y--
			case DV, V:
				x--
			}
			pos--
		}
		return matchedPos
	}
	pos := m - 2
	shortPos := n - 2
	for x != 0 || y != 0 {
		if pos < 0 {
			break
		}
		switch cmat[x][y] {
		case D:
			x--
			y--
			if s2[pos] == s1[shortPos] {
				matchedPos = append(matchedPos, pos)
				shortPos--
			}
		case DH, H:
			y--
		case DHV, HV, DV, V:
			x--
		}
		pos--
	}

	return matchedPos
}
