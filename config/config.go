package config

const (
	DEBUG = false
	// SCORE = "NeedlemanWunsch"
	SCORE          = "SmithWaterman"
	HighlightBegin = "\x1b[38;5;198m"
	HighlightEnd   = "\x1b[0m"
	GAP            = 2
	MATCH          = 2
	MISMATCH       = -1
)
