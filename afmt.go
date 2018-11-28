package afmt

import (
// "./internal/printer"
)

type testStruct struct {
	Level0101 string
	Level0102 struct {
		Level0201 string
		Level0202 int
		Level0203 struct {
			Level0301 string
		}
	}
}

// PrintTree ...
func PrintTree(structure interface{}) {

	// printer.PrintStruct(structure)

}
