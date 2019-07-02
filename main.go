package main

import "github.com/naronA/fuzzyfinder/score"

func main() {
	score.NeedlemanWunsch("ATACABACCCC", "ACBCCCC", true)
	score.NeedlemanWunsch("AC", "ATAC", true)
}
