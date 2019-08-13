package score

import (
	"sort"
)

type Range struct {
	Start int
	End   int
}

type Ranges []*Range

func (r Ranges) Len() int {
	return len(r)
}

func (r Ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Ranges) Less(i, j int) bool {
	return r[i].Start < r[j].Start
}

func overlap(start, end int, m Range) bool {
	if m.Start < start && start < m.End && end >= m.End {
		return true
	}
	return false
}

func merge(r Ranges, newR *Range) Ranges {
	if len(r) == 0 {
		return Ranges{newR}
	}

	newRanges := Ranges{}
	for _, m := range r {
		switch {
		case newR.Start >= m.Start && m.End >= newR.End:
			// m       |----------|
			// newR       |----|
			// Result  |----------|
			newR.Start = m.Start
			newR.End = m.End
		case newR.Start <= m.Start && m.Start <= newR.End && m.End >= newR.End:
			// m           |-------|
			// newR    |-----|
			// Result  |----------|
			newR.End = m.End
		case m.Start <= newR.Start && newR.Start <= m.End && m.End <= newR.End:
			// m       |------|
			// newR        |------|
			// Result  |----------|
			newR.Start = m.Start
		case newR.Start <= m.Start && m.Start <= newR.End && newR.Start <= m.End && m.End <= newR.End:
			// m           |-------|
			// newR    |---------------|
			// Result  |---------------| do nothing
		default:
			newRanges = append(newRanges, m)
		}
	}
	newRanges = append(newRanges, newR)
	return newRanges
}

type Finder struct {
	Source string
	Inputs []string
}

func (f Finder) Score() int {
	var sc int
	for _, input := range f.Inputs {
		score, _ := CalcScore(f.Source, input)
		sc += score
	}
	return sc
}

func (f Finder) Matches() Ranges {
	matches := Ranges{}
	for _, input := range f.Inputs {
		starts := IndicesAll(f.Source, input)
		for _, start := range starts {
			end := start + len(input)
			r := &Range{Start: start, End: end}
			matches = merge(matches, r)
			// matches = mergeRange(matches, m)
		}
	}
	sort.Sort(matches)
	return matches
}

// マッチした文字列をハイライトするために、対象文字の前後に制御文字を埋め込む
func (f Finder) String() string {
	hBegin := []rune("\x1b[38;5;198m")
	hEnd := []rune("\x1b[0m")
	source := []rune(f.Source)
	headStart := 0
	matches := f.Matches()
	highlighted := make([]rune, 0, len(source)+len(matches)*(len(hBegin)+len(hEnd)))
	for i, m := range matches {
		head := source[headStart:m.Start]
		term := source[m.Start:m.End]
		highlighted = append(highlighted, head...)
		highlighted = append(highlighted, hBegin...)
		highlighted = append(highlighted, term...)
		highlighted = append(highlighted, hEnd...)

		if i+1 < len(matches) {
			headStart = m.End
		} else {
			tail := source[m.End:]
			highlighted = append(highlighted, tail...)
		}
	}
	return string(highlighted)
}

type Finders []Finder

func (f Finders) Len() int {
	return len(f)
}

func (f Finders) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Finders) Less(i, j int) bool {
	return f[i].Score() < f[j].Score()
}
