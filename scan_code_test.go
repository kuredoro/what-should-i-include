package wsii_test

import (
    "testing"

    "github.com/kuredoro/what-should-i-include"
)

func TestScanCode(t *testing.T) {
    t.Run("empty code", func(t *testing.T) {
        code := ""

        got := wsii.ScanCode(code)
        if got != nil {
            t.Errorf("got includes %v, want none", got)
        }
    })
}
