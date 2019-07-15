package score

import (
	"fmt"
	"testing"
)

func TestIndicesAll(t *testing.T) {
	ids := IndicesAll("abcdddabcdddabc", "abc")
	fmt.Println(ids)
	if ids[0] != 0 || ids[1] != 6 || ids[2] != 12 || len(ids) > 3 {
		t.Fail()
	}
}

func TestIndicesAll2(t *testing.T) {
	ids := IndicesAll("score/score.go", "or")
	fmt.Println(ids)
	if ids[0] != 2 || ids[1] != 8 || len(ids) > 2 {
		t.Fail()
	}

}
