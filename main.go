package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/naronA/fuzzyfinder/score"
)

func main() {
	flag.Parse()
	args := flag.Args()

	input := args[0]
	paths := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			switch info.Name() {
			case ".git":
				return filepath.SkipDir
			case ".mypy_cache":
				return filepath.SkipDir
			}
		}
		if err != nil {
			return err
		}
		if info.Name() != "." {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	finders := Finders{}
	for _, src := range paths {
		sc := score.NeedlemanWunsch(src, input)
		if strings.Contains(src, input) {
			f := Finder{Score: sc, Source: src, Input: input}
			finders = append(finders, f)
		}
	}
	sort.Sort(sort.Reverse(finders))
	for _, f := range finders {
		fmt.Println(f)
	}
}

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
	// return len(f[i].Pointers) < len(f[j].Pointers) || f[i].Score < f[j].Score
	return f[i].Score < f[j].Score
}

func (f Finder) String() string {
	highligh := []rune{}
	source := []rune(f.Source)
	input := []rune(f.Input)
	index := strings.Index(f.Source, f.Input)
	for i, c := range source {
		if i == index {
			highligh = append(highligh, []rune("\x1b[38;5;198m")...)
		}
		highligh = append(highligh, c)
		if i == index+len(input)-1 {
			highligh = append(highligh, []rune("\x1b[0m")...)
		}

	}
	return string(highligh)
}
