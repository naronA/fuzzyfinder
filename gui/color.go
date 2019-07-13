package gui

import "sort"

func HighlightRune(str []rune, p []int) []rune {
	highlight := []rune{}
	sort.Ints(p)
	for i, c := range str {
		if len(p) > 0 && i == p[0] {
			p = p[1:]
			highlight = append(highlight, []rune("\x1b[38;5;198m")...)
			highlight = append(highlight, c)
			highlight = append(highlight, []rune("\x1b[0m")...)
		} else {
			highlight = append(highlight, c)
		}
	}
	return highlight

}
