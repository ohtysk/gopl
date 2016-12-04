package main

import (
	"testing"
)

func TestMyjoin(t *testing.T) {
	args := []string{"abc", "hoge", "ZSDF", "1234"}
	expected := "abc hoge ZSDF 1234"
	actual := myjoin(args)
	if expected != actual {
		t.Errorf("error expectd %s but %s", expected, actual)
	}
}
