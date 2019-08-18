package score

import (
	"github.com/naronA/fuzzyfinder/config"
)

func CalcScore(path, input string) int {
	switch config.SCORE {
	case "NeedlemanWunsch":
		return NeedlemanWunsch(path, input)
	case "LevenshteinDistance":
		return LevenshteinDistance(path, input)
	case "SmithWaterman":
		return SmithWaterman(path, input)

	}
	return NeedlemanWunsch(path, input)
}
