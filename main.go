package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/naronA/fuzzyfinder/score"
)

func main() {
	ignore := []string{".git", ".mypy_cache"}
	flag.Parse()
	args := flag.Args()

	input := args[0]

	// finders := score.Finders{}
	wg := &sync.WaitGroup{}
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
			wg.Add(1)
			go func() {
				if strings.Contains(path, input) {
					sc := score.NeedlemanWunsch(path, input)
					f := score.Finder{Score: sc, Source: path, Input: input}
					fmt.Println(f)
					// finders = append(finders, f)
				}
				wg.Done()
			}()
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	wg.Wait()
	// sort.Sort(sort.Reverse(finders))
	// for _, f := range finders {
	// 	fmt.Println(f)
	// }
}
