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
	ignore := []string{".git", ".mypy_cache"}
	flag.Parse()
	args := flag.Args()

	input := args[0]

	finders := score.Finders{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			for _, dir := range ignore {
				if info.Name() == dir {
					return filepath.SkipDir
				}
			}
		}
		if info.Name() != "." {
			if strings.Contains(path, input) {
				sc := score.NeedlemanWunsch(path, input)
				f := score.Finder{Score: sc, Source: path, Input: input}
				finders = append(finders, f)
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	sort.Sort(sort.Reverse(finders))
	for _, f := range finders {
		fmt.Println(f)
	}
}
