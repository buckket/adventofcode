package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		filename string
		expected string
	}{
		{"test_input", "95437"},
		{"real_input", "919137"},
	}

	for _, e := range tests {
		f, _ := os.Open(e.filename)
		defer f.Close()

		result := Part1(f)
		if result != e.expected {
			t.Errorf("Result is incorrect, got: %s, want: %s.", result, e.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		filename string
		expected string
	}{
		{"test_input", "24933642"},
		{"real_input", "2877389"},
	}

	for _, e := range tests {
		f, _ := os.Open(e.filename)
		defer f.Close()

		result := Part2(f)
		if result != e.expected {
			t.Errorf("Result is incorrect, got: %s, want: %s.", result, e.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	f, _ := os.Open("real_input")
	defer f.Close()

	for n := 0; n < b.N; n++ {
		Part1(f)
		f.Seek(0, 0)
	}
}

func BenchmarkPart2(b *testing.B) {
	f, _ := os.Open("real_input")
	defer f.Close()

	for n := 0; n < b.N; n++ {
		Part2(f)
		f.Seek(0, 0)
	}
}
