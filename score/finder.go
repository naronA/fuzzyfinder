package score

import "strings"

type Finder struct {
	Score  int
	Source string
	Input  string
}

type Finders []Finder

func (f Finders) Len() int {
	return len(f)
}

func (f Finders) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Finders) Less(i, j int) bool {
	return f[i].Score < f[j].Score
}

func (f Finder) String() string {
	index := strings.Index(f.Source, f.Input)
	if index == -1 {
		return f.Source
	}
	hBegin := "\x1b[38;5;198m"
	hEnd := "\x1b[0m"
	source := []rune(f.Source)
	input := []rune(f.Input)
	highligh := make([]rune, 0, len(hBegin)+len(hEnd)+len(source))
	highligh = append(highligh, source[:index]...)
	highligh = append(highligh, []rune(hBegin)...)
	highligh = append(highligh, source[index:index+len(input)]...)
	highligh = append(highligh, []rune(hEnd)...)
	highligh = append(highligh, source[index+len(input):]...)
	return string(highligh)
}
