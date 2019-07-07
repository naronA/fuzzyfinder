package score

import "testing"

func TestNeedlemanWunsch01(t *testing.T) {
	NeedlemanWunsch("ATACABACCCC", "ACBCCCC", true)
}

func TestNeedlemanWunsch02(t *testing.T) {
	NeedlemanWunsch("AC", "ATAC", true)
}
