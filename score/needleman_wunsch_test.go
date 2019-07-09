package score

import "testing"

func TestNeedlemanWunsch01(t *testing.T) {
	NeedlemanWunsch("ATACABACCCC", "ACBCCCC")
}

func TestNeedlemanWunsch02(t *testing.T) {
	NeedlemanWunsch("AC", "ATAC")
}
