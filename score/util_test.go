package score

import (
	"fmt"
	"testing"
)

func TestIndicesAll(t *testing.T) {
	ids := IndicesAll("abcdddabcdddabc", "abc")
	if ids[0] != 0 || ids[1] != 6 || ids[2] != 12 || len(ids) > 3 {
		fmt.Println(ids)
		t.Fail()
	}
}

func TestIndicesAll2(t *testing.T) {
	ids := IndicesAll("score/score.go", "or")
	if ids[0] != 2 || ids[1] != 8 || len(ids) > 2 {
		fmt.Println(ids)
		t.Fail()
	}

}
