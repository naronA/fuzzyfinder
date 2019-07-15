package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/naronA/fuzzyfinder/score"
)

func main() {
	start := time.Now()
	ignore := []string{
		// ".git",
		// ".mypy_cache",
		// ".vscode",
		// ".idea",
		// "node_modules",
		// "vendor",
	}
	flag.Parse()
	args := flag.Args()
	finders := score.Finders{}

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
			func(path string) {
				defer wg.Done()
				for _, input := range args {
					if !strings.Contains(path, input) {
						return
					}
				}
				f := score.Finder{Source: path, Inputs: args}
				finders = append(finders, f)
			}(path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	wg.Wait()
	for _, f := range finders {
		fmt.Println(f)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
