// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

import (
	"testing"
)

func TestBasicString(t *testing.T) {
	var err error
	var result string

	predictedResult := "Lorem ipsum dolor sit amet\n"

	tp := NewTreePrinter()

	var testValue = "Lorem ipsum dolor sit amet"

	if result, err = tp.Print(testValue); err != nil {
		t.Errorf("Error fired! %s", err.Error())
	}

	if result != predictedResult {
		t.Errorf("Unpredictable result")
	}
}

func TestBasicInteger(t *testing.T) {
	var err error
	var result string

	predictedResult := "10\n"

	tp := NewTreePrinter()

	var testValue = 10

	if result, err = tp.Print(testValue); err != nil {
		t.Errorf("Error fired! %s", err.Error())
	}

	if result != predictedResult {
		t.Errorf("Unpredictable result")
	}
}

func TestBasicBoolean(t *testing.T) {
	var err error
	var result string

	predictedResult := "true\n"

	tp := NewTreePrinter()

	var testValue = true

	if result, err = tp.Print(testValue); err != nil {
		t.Errorf("Error fired! %s", err.Error())
	}

	if result != predictedResult {
		t.Errorf("Unpredictable result")
	}
}

type testBasicStruct struct {
	Item1 string
	Item2 int
	Item3 *string
}

// func TestBasicStruct(t *testing.T) {
// 	var err error
// 	var result string

// 	predictedResult := "testBasicStruct:\n"
// 	predictedResult += "├── Item1: Lorem ipsum dolor sit amet\n"
// 	predictedResult += "├── Item2: 5\n"
// 	predictedResult += "└── Item3: Lorem ipsum dolor sit amet\n"

// 	tp := NewTreePrinter()
// 	testString := "Lorem ipsum dolor sit amet"

// 	var testValue = testBasicStruct{
// 		Item1: testString,
// 		Item2: 5,
// 		Item3: &testString,
// 	}

// 	if result, err = tp.Print(testValue); err != nil {
// 		t.Errorf("Error fired! %s", err.Error())
// 	}

// 	if result != predictedResult {
// 		t.Logf(result)
// 		t.Errorf("Unpredictable result")
// 	}
// }
