package main

import "testing"

func TestJoinwithfor(t *testing.T) {
	args := []string{"abc", "hoge", "ZSDF", "1234"}
	expected := "hoge ZSDF 1234"
	actual := joinwithfor(args)
	if expected != actual {
		t.Errorf("error expectd %s but %s", expected, actual)
	}
}

func TestJoinofstrings(t *testing.T) {
	args := []string{"abc", "hoge", "ZSDF", "1234"}
	expected := "hoge ZSDF 1234"
	actual := joinofstrings(args)
	if expected != actual {
		t.Errorf("error expectd %s but %s", expected, actual)
	}
}

func BenchmarkJoinwithfor(b *testing.B) {
	args := []string{"abc", "hoge", "ZSDF", "1234"}
	for i := 0; i < b.N; i++ {
		joinwithfor(args)
	}
}

func BenchmarkJoinofstrings(b *testing.B) {
	args := []string{"abc", "hoge", "ZSDF", "1234"}
	for i := 0; i < b.N; i++ {
		joinofstrings(args)
	}
}
