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

type Finder struct {
	Source string
	Inputs []string
}

func (f Finder) Score() int {
	var sc int
	for _, input := range f.Inputs {
		sc += CalcScore(f.Source, input)
	}
	return sc
}

func overlap(start, end int, m Range) bool {
	if m.Start < start && start < m.End && end >= m.End {
		return true
	}
	return false
}

func (f Finder) Matches() Ranges {
	matches := Ranges{}
	for _, input := range f.Inputs {
		starts := IndicesAll(f.Source, input)
		for _, start := range starts {
			end := start + len(input)
			isOverlap := false
			for _, m := range matches {
				if m.Start <= start && start <= m.End && end > m.End {
					isOverlap = true
					m.End = end
					break
				} else if start < m.Start && m.Start <= end && end <= m.End {
					isOverlap = true
					m.Start = start
					break
				} else if m.Start <= start && start <= m.End && m.Start <= end && end <= m.End {
					isOverlap = true
					break
				}
			}
			if isOverlap == false {
				m := &Range{Start: start, End: end}
				matches = append(matches, m)
			}
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
