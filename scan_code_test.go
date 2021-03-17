package wsii_test

import (
	"testing"

	"github.com/kuredoro/what-should-i-include"
)

func TestScanCode(t *testing.T) {
    t.Run("empty code", func(t *testing.T) {
        code := ""

        got := wsii.ScanCode(code)

        wsii.AssertIncludes(t, got, nil)
    })

    t.Run("one double-quote include", func(t *testing.T) {
        code := "#include \"test\""

        got := wsii.ScanCode(code)
        want := []string{"test"}

        wsii.AssertIncludes(t, got, want)
    })

    /*
    t.Run("one angle-bracket include", func(t *testing.T) {
        code := "#include <foo>"

        got := wsii.ScanCode(code)
        want := []string{"foo"}

        wsii.AssertIncludes(t, got, want)
    })
    */
}
