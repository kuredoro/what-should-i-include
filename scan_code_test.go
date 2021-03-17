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

func TestScanCode(t *testing.T) {
    t.Run("empty file", func(t *testing.T) {
        code := ""

        got := wsii.ScanCode(code)

        wsii.AssertIncludes(t, got, nil)
    })

    t.Run("hello world", func(t *testing.T) {
        code := `#include <iostream>
using namespace std;

int main() {
    cout << "Hello, world!" << endl;
    return 0;
}`

        got := wsii.ScanCode(code)

        want := []string{"iostream"}

        wsii.AssertIncludes(t, got, want)
    })

    t.Run("lots of includes", func(t *testing.T) {
        code := `
#include <iostream>
#include <functional>

#include <boost/lexical_cast.hpp>
#include <boost/optional.hpp>

#include "myheader.hpp"

#if __has_include(<wsii.tweaks.h>)
    #include <wsii.tweaks.h>
#endif

const char * msg = "Hello, world!"

int main() {
    std::cout << msg << std::endl;
    return 0;
}`

        got := wsii.ScanCode(code)

        want := []string{
            "iostream",
            "functional",
            "boost/lexical_cast.hpp",
            "boost/optional.hpp",
            "myheader.hpp",
            "wsii.tweaks.h",
        }

        wsii.AssertIncludes(t, got, want)
    })
}
