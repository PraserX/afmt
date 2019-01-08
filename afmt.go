package afmt

import (
	"fmt"
)

// PrintTree ...
func PrintTree(structure interface{}) error {
	var err error
	var tree string

	p := NewTreePrinter()
	if tree, err = p.Print(structure); err == nil {
		fmt.Printf(tree)
	}

	return err
}
