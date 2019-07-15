package score

import (
	"github.com/naronA/fuzzyfinder/config"
)

func CalcScore(path, input string) (int, []int) {
	switch config.SCORE {
	case "NeedlemanWunsch":
		return NeedlemanWunsch(path, input)
	case "LevenshteinDistance":
		return LevenshteinDistance(path, input)
	}
	return NeedlemanWunsch(path, input)
}
