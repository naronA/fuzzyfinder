package main

import "github.com/naronA/fuzzyfinder/score"

func main() {
	score.NeedlemanWunsch([]rune("ATACABACCCC"), []rune("ACBCCCC"), true)
	score.NeedlemanWunsch([]rune("AC"), []rune("ATAC"), true)
}
