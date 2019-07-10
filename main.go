package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/naronA/fuzzyfinder/score"
)

func main() {
	// gui.Gui()
	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }

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
	input := "util"
	finders := Finders{}
	for _, str := range paths {
		sc, matched := score.NeedlemanWunsch(str, input)
		if len(matched) > 0 {
			f := Finder{Score: sc, Str: str, Pointers: matched}
			finders = append(finders, f)
		}
	}
	sort.Sort(sort.Reverse(finders))
	fmt.Println(finders)
}

type Finder struct {
	Score    int
	Str      string
	Pointers []int
}

type Finders []Finder

func (f Finders) Len() int {
	return len(f)
}

func (f Finders) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Finders) Less(i, j int) bool {
	return len(f[i].Pointers) < len(f[j].Pointers) || f[i].Score < f[j].Score
}

func (f Finder) String() string {
	highlighted := []rune{}
	original := []rune(f.Str)
	sort.Ints(f.Pointers)
	for i, c := range original {
		if len(f.Pointers) > 0 && i == f.Pointers[0] {
			f.Pointers = f.Pointers[1:]
			highlighted = append(highlighted, []rune("\x1b[38;5;127m")...)
			highlighted = append(highlighted, c)
			highlighted = append(highlighted, []rune("\x1b[0m")...)
		} else {
			highlighted = append(highlighted, c)
		}
	}
	return string(highlighted)
}
