package score

import (
	"sort"
)

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

func overlap(start, end int, m Range) bool {
	if m.Start < start && start < m.End && end >= m.End {
		return true
	}
	return false
}

func mergeRange(matches Ranges, newR *Range) Ranges {
	if len(matches) == 0 {
		return Ranges{newR}
	}

	newRanges := Ranges{}
	for _, m := range matches {
		if newR.Start >= m.Start && m.End >= newR.End {
			// m    |----------|
			// newR    |----|
			newR.Start = m.Start
			newR.End = m.End
		} else if newR.Start <= m.Start && m.Start <= newR.End && m.End >= newR.End {
			// m       |-------|
			// newR |-----|
			newR.End = m.End
		} else if m.Start <= newR.Start && newR.Start <= m.End && m.End <= newR.End {
			// m     |-------|
			// newR       |------|
			newR.Start = m.Start
		} else if newR.Start <= m.Start && m.Start <= newR.End && newR.Start <= m.End && m.End <= newR.End {
			// m           |-------|
			// newR |---------------|
			// do nothing
		} else {
			newRanges = append(newRanges, m)
		}
	}
	newRanges = append(newRanges, newR)
	return newRanges
}

func (f Finder) Matches() Ranges {
	matches := Ranges{}
	for _, input := range f.Inputs {
		starts := IndicesAll(f.Source, input)
		for _, start := range starts {
			end := start + len(input)
			m := &Range{Start: start, End: end}
			matches = mergeRange(matches, m)
		}
	}
	sort.Sort(matches)
	return matches
}

func (f Finder) String() string {
	hBegin := []rune("\x1b[38;5;198m")
	hEnd := []rune("\x1b[0m")
	highligh := []rune(f.Source)
	for i, m := range f.Matches() {
		gap := i * (len(hBegin) + len(hEnd))
		head := highligh[:m.Start+gap]
		term := highligh[m.Start+gap : m.End+gap]
		tail := highligh[m.End+gap:]

		highligh = make([]rune, 0, len(highligh)+(len(hBegin)+len(hEnd)))
		highligh = append(highligh, head...)
		highligh = append(highligh, hBegin...)
		highligh = append(highligh, term...)
		highligh = append(highligh, hEnd...)
		highligh = append(highligh, tail...)
	}

	return string(highligh)
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
