package main

import (
	"testing"
)

func TestExtract(t *testing.T) {

	text := []string{
		"かきくけこ",
		"あいう",
		"あいうえお",
		"ABC",
		"123",
	}

	runes := []rune("123ABCあいうえおかきくけこ")

	r := extract(text)

	for i, c := range r {
		if c != runes[i] {
			t.Errorf("got %v\nwant %v", c, runes[i])
		}
	}

}
