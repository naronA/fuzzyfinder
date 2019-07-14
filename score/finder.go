package score

import (
	"strings"
)

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

func (f Finder) String() string {
	hBegin := "\x1b[38;5;198m"
	hEnd := "\x1b[0m"
	highligh := []rune(f.Source)
	for _, input := range f.Inputs {
		point := strings.Index(string(highligh), input)
		head := highligh[:point]
		term := highligh[point : point+len(input)]
		tail := highligh[point+len(input):]

		highligh = make([]rune, 0, len(highligh)+2)
		highligh = append(highligh, head...)
		highligh = append(highligh, '\t')
		highligh = append(highligh, term...)
		highligh = append(highligh, '\v')
		highligh = append(highligh, tail...)
	}

	result := string(highligh)
	for strings.Contains(result, "\t") || strings.Contains(result, "\v") {
		result = strings.Replace(result, "\t", hBegin, 1)

		// タグが入れ子になっていたら入れ子の部分を削除する
		// \t ab \t cA \v BC \v  => \t ab cA BC \v
		for strings.Contains(result, "\t") && strings.Index(result, "\t") < strings.Index(result, "\v") {
			result = strings.Replace(result, "\t", "", 1)
			result = strings.Replace(result, "\v", "", 1)
		}

		result = strings.Replace(result, "\v", hEnd, 1)
	}
	return result
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
