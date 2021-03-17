package wsii_test

import (
	"reflect"
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

    t.Run("one double-quote include", func(t *testing.T) {
        code := "#include \"test\""

        got := wsii.ScanCode(code)
        want := []string{"test"}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got includes %v, want %v", got, want)
        }
    })
}
