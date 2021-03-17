package wsii_test

import (
	"testing"

	"github.com/kuredoro/what-should-i-include"
)

func assertString(t *testing.T, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got string %q, want %q", got, want)
    }
}

func TestScanLine(t *testing.T) {

    cases := []struct{
        title string
        input string
        want string
    }{
        {"empty line", "", ""},
        {"one double-quote include", "#include \"test\"", "test"},
        {"one angle-bracket include",  "#include <foo>", "foo"},
        {"string literal definition",  "const char * foo = \"bar\"", ""},
        {"an ambiguous if statement",  "if (0 < a && b > 0)", ""},
        {"space-idented",  "    #include <hello>", "hello"},
        {"tab-idented",  "\t#include <hello>", "hello"},
    }

    for _, test := range cases {
        t.Run(test.title, func(t *testing.T) {
            got := wsii.ScanLine(test.input)

            assertString(t, got, test.want)

        })
    }
}

