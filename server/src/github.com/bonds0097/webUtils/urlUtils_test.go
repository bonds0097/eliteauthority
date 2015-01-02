package webUtils

import (
	"testing"
)

type TestCase struct {
	in  string
	out string
}

var GenerateSlugTests = []TestCase{
	{"Test Station", "test-station"},
	{"Harry's Hub", "harrys-hub"},
	{"()()()()", ""},
	{"Bob's Place (Perhaps)", "bobs-place-perhaps"},
	{"X-1!~Bravo Man", "x-1~bravo-man"},
	{"HELLO_THERE", "hello_there"},
}

var NormalizeWhitespaceTests = []TestCase{
	{"test case", "test case"},
	{" test case", "test case"},
	{"test case ", "test case"},
	{"  test  case  ", "test case"},
	{"  test  case  ", "test case"},
}

func TestGenerateSlug(t *testing.T) {
	for _, testCase := range GenerateSlugTests {
		out, _ := GenerateSlug(testCase.in)
		if out != testCase.out {
			t.Error("For", testCase.in,
				"expected", testCase.out,
				"got", out)
		}
	}
}

func TestNormalizeWhitesoace(t *testing.T) {
	for _, testCase := range NormalizeWhitespaceTests {
		out := NormalizeWhitespace(testCase.in)
		if out != testCase.out {
			t.Error("For", testCase.in,
				"expected", testCase.out,
				"got", out)
		}
	}
}
