// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

import (
	"testing"
	"fmt"
)

func TestBasic(t *testing.T) {
	cp := NewColPrinter()

	for _, c := range []struct {
		in   string
		want string
		param int
	}{
		{"Lorem ipsum dolor sit amet", "Lorem\nipsum\ndolor\nsit\namet", 5},
		{"Lorem ipsum dolor sit amet", "Lorem ipsum\ndolor sit amet", 15},
	} {
		if cp.Print(c.param, c.in) != c.want {
			t.Errorf("Unpredictable result for column count: %d", c.param)
		}
	}
}

func TestExtended(t *testing.T) {
	cp := NewColPrinter()

	for _, c := range []struct {
		in1   string
		in2 string
		want string
		param int
	}{
		{"Lorem ipsum dolor sit amet", "Lorem ipsum dolor sit amet", "Lorem\nipsum\ndolor\nsit\namet\nLorem\nipsum\ndolor\nsit\namet", 5},
		{"Lorem ipsum dolor sit amet", "Lorem ipsum dolor sit amet", "Lorem ipsum\ndolor sit amet\nLorem ipsum\ndolor sit amet", 15},
	} {
		if cp.Print(c.param, c.in1, c.in2) != c.want {
			fmt.Println(cp.Print(c.param, c.in1, c.in2))
			t.Errorf("Unpredictable result for column count: %d", c.param)
		}
	}
}