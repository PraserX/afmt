// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

import (
	"fmt"
	"io"
	"os"
)

// PrintTree print input to standard output structured to tree representation.
func PrintTree(structure interface{}) error {
	var err error
	var tree string

	p := NewTreePrinter()
	if tree, err = p.Print(structure); err == nil {
		fmt.Printf(tree)
	}

	return err
}

// FprintCol formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string. Newline is always
// added in order to respect column maximum characters. It returns the number of
// bytes written and any write error encountered.
func FprintCol(w io.Writer, lineLimit int, a ...interface{}) (n int, err error) {
	p := NewColPrinter()
	str := p.Print(lineLimit, a...)
	n, err = fmt.Fprint(w, str)
	return n, err
}

// FprintlnCol formats using the default formats for its operands and writes to
// w. Spaces are always added between operands and a newline is appended.
// Newline is always added in order to respect column maximum characters. It
// returns the number of bytes written and any write error encountered.
func FprintlnCol(w io.Writer, lineLimit int, a ...interface{}) (n int, err error) {
	p := NewColPrinter()
	str := p.Print(lineLimit, a...)
	n, err = fmt.Fprintln(w, str)
	return n, err
}

// PrintCol formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a string.
// Newline is always added in order to respect column maximum characters. It
// returns the number of bytes written and any write error encountered.
func PrintCol(lineLimit int, a ...interface{}) (n int, err error) {
	p := NewColPrinter()
	str := p.Print(lineLimit, a...)
	n, err = fmt.Fprint(os.Stdout, str)
	return n, err
}

// PrintlnCol formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. Newline is always added in order to respect column maximum
// characters. It returns the number of bytes written and any write error
// encountered.
func PrintlnCol(lineLimit int, a ...interface{}) (n int, err error) {
	p := NewColPrinter()
	str := p.Print(lineLimit, a...)
	n, err = fmt.Fprintln(os.Stdout, str)
	return n, err
}
