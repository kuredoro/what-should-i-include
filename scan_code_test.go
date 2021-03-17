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
    t.Run("empty line", func(t *testing.T) {
        line := ""

        got := wsii.ScanLine(line)

        assertString(t, got, "")
    })

    t.Run("one double-quote include", func(t *testing.T) {
        line := "#include \"test\""

        got := wsii.ScanLine(line)
        want := "test"

        assertString(t, got, want)
    })

    t.Run("string literal definition", func (t *testing.T) {
        line := "const char * foo = \"bar\""

        got := wsii.ScanLine(line)

        assertString(t, got, "")
    })

    t.Run("an ambiguous if statement", func (t *testing.T) {
        line := "if (0 < a && b > 0)"

        got := wsii.ScanLine(line)

        assertString(t, got, "")
    })

    t.Run("one angle-bracket include", func(t *testing.T) {
        line := "#include <foo>"

        got := wsii.ScanLine(line)
        want := "foo"

        assertString(t, got, want)
    })
}

