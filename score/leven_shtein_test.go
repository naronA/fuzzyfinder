package main

import (
	"fmt"
	"testing"
)

type ArgResult struct {
	Str1   string
	Str2   string
	Result int
}

func TestLevenshteinDistance0(t *testing.T) {
	actual := LevenshteinDistance([]rune("aaaaa"), []rune("aaaaa"))
	expected := 0
	if expected != actual {
		fmt.Println(expected, actual)
		t.Fail()
	}
}

func TestLevenshteinDistance1(t *testing.T) {
	actual := LevenshteinDistance([]rune("aaaaa"), []rune("baaaa"))
	expected := 1
	if expected != actual {
		fmt.Println(expected, actual)
		t.Fail()
	}
}

func TestLevenshteinDistance2(t *testing.T) {
	actual := LevenshteinDistance([]rune("aaaaa"), []rune("baaba"))
	expected := 2
	if expected != actual {
		fmt.Println(expected, actual)
		t.Fail()
	}
}

func TestLevenshteinDistanceDifferenceLength(t *testing.T) {
	actual := LevenshteinDistance([]rune("aaaaa"), []rune("aaaaabbb"))
	expected := 3
	if expected != actual {
		fmt.Println(expected, actual)
		t.Fail()
	}
}
