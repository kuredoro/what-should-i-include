package wsii

import (
    "testing"
)

func TestSubstrEnclosedBy(t *testing.T) {
    t.Run("empty", func(t *testing.T) {
        got := substrEnclosedBy("", '(', ')')

        assertString(t, got, "")
    })

    t.Run("only braces", func(t *testing.T) {
        got := substrEnclosedBy("{}", '{', '}')

        assertString(t, got, "")
    })

    t.Run("braces at beginning and end", func(t *testing.T) {
        got := substrEnclosedBy("{foo}", '{', '}')

        assertString(t, got, "foo")
    })

    t.Run("braces somewhere inside", func(t *testing.T) {
        got := substrEnclosedBy("this is {foo}. It's not bar.", '{', '}')

        assertString(t, got, "foo")
    })

    t.Run("no closing", func(t *testing.T) {
        got := substrEnclosedBy("{foo is bar", '{', '}')

        assertString(t, got, "{foo is bar")
    })

    t.Run("no openning", func(t *testing.T) {
        got := substrEnclosedBy("or is it}", '{', '}')

        assertString(t, got, "or is it}")
    })

    t.Run("switched places", func(t *testing.T) {
        got := substrEnclosedBy("gotch} ya { good", '{', '}')

        assertString(t, got, "gotch} ya { good")
    })

    t.Run("single quote", func(t *testing.T) {
        got := substrEnclosedBy("a single ' quote", '\'', '\'')

        assertString(t, got, "a single ' quote")
    })

    t.Run("quote", func(t *testing.T) {
        got := substrEnclosedBy("\"Think twice, do once\"", '"', '"')

        assertString(t, got, "Think twice, do once")
    })

    t.Run("nested brackets", func(t *testing.T) {
        got := substrEnclosedBy("a[b[5] + 2]", '[', ']')

        assertString(t, got, "b[5] + 2")
    })

    t.Run("several pairs", func(t *testing.T) {
        got := substrEnclosedBy("(1 + 2) - (3 + 4)", '(', ')')

        assertString(t, got, "1 + 2) - (3 + 4")
    })
}
