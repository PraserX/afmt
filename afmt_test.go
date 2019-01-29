// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

import (
	"testing"
)

type testStruct01 struct {
	Level0101 testStruct01x2
	Level0102 struct {
		Level0201 string
		Level0202 int
		Level0203 struct {
			Level0301 string
		}
	}
	Level0103 bool
	Level0104 int
	Level0105 []string
}

type testStruct02 struct {
	Level0101 *testStruct01x2
	Level0102 struct {
		Level0201 *string
		Level0202 *int
	}
	Level0103 *testStruct01x2
}

type testStruct01x2 struct {
	Level0201 string
}

func TestPrintTree01(t *testing.T) {
	str := testStruct01{}
	str.Level0101.Level0201 = "Lorem ipsum dolor sit amet"
	str.Level0102.Level0201 = "Lorem ipsum dolor sit amet"
	str.Level0102.Level0202 = 10
	str.Level0102.Level0203.Level0301 = "Lorem ipsum dolor sit amet"
	str.Level0103 = false
	str.Level0104 = 1458
	str.Level0105 = []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
	PrintTree(str)

	// t.FailNow()
}

func TestPrintTree02(t *testing.T) {
	var strct *testStruct01x2
	var strng string
	var intgr int

	strct = &testStruct01x2{Level0201: "Lorem ipsum dolor sit amet"}
	strng = "Lorem ipsum dolor sit amet"
	intgr = 10

	str := testStruct02{}
	str.Level0101 = strct
	str.Level0102.Level0201 = &strng
	str.Level0102.Level0202 = &intgr
	str.Level0103 = nil
	PrintTree(str)

	// t.FailNow()
}
