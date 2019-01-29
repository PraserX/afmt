// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

import (
	"fmt"
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
