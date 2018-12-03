package afmt

// PrintTree ...
func PrintTree(structure interface{}) {
	p := NewPrinter()
	p.Tree(structure)
}
