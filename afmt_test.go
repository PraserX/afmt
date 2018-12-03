package afmt

import (
	"testing"
)

type testStruct struct {
	Level0101 testStruct2
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

type testStruct2 struct {
	Level0201 string
}

func TestPrintTree(t *testing.T) {

}

func TestEchoTree(t *testing.T) {
	str := testStruct{}
	str.Level0101.Level0201 = "Lorem ipsum dolor sit amet"
	str.Level0102.Level0201 = "Lorem ipsum dolor sit amet"
	str.Level0102.Level0202 = 10
	str.Level0102.Level0203.Level0301 = "Lorem ipsum dolor sit amet"
	str.Level0103 = false
	str.Level0104 = 1458
	str.Level0105 = []string{"Lorem", "ipsum", "dolor", "sit", "amet"}

	PrintTree(str)
}
