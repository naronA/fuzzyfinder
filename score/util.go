package score

import (
	"strings"
	"unicode/utf8"
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
				matchedPos = append(matchedPos, shortPos)
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

func sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func last(nums []int) int {
	last := 0
	for _, val := range nums {
		last = val
	}
	return last
}

func IndicesAll(src, substr string) []int {
	indices := []int{}

	slice := []rune(src)
	for i := 0; strings.Contains(string(slice), substr); i++ {
		idx := strings.Index(string(slice), substr)
		count := utf8.RuneCountInString(string(slice)[:idx])
		substrCount := utf8.RuneCountInString(substr)
		slice = slice[count+substrCount:]
		newIdx := func() int {
			if i >= 1 {
				return count + last(indices) + substrCount
			}
			return count + last(indices)
		}()
		indices = append(indices, newIdx)
	}
	return indices
}
