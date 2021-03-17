package wsii

import (
    "strings"
)

func substrEnclosedBy(str string, opening, closing rune) string {
    startPos := strings.IndexRune(str, opening)

    if startPos == -1 {
        return str 
    }

    endPos := strings.LastIndexFunc(str, func(r rune) bool {
        return r == closing
    })

    if endPos == -1 {
        return str
    }

    if startPos >= endPos {
        return str
    }

    return str[startPos+1:endPos]
}

