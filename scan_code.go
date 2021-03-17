package wsii

import (
    "strings"
)

func ScanCode(code string) []string {
    parts := strings.Split(code, "\"")

    if code != "" && len(parts) == 1 {
        head := strings.Split(code, "<")
        parts = head[:1]
        parts = append(parts, strings.Split(head[1], ">")...)
    }

    if len(parts) == 3 {
        return []string{parts[1]}
    }

    return nil
}
