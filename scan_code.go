package wsii

import (
    "strings"
)

func ScanCode(code string) []string {
    parts := strings.Split(code, "\"")

    if len(parts) == 3 {
        return []string{parts[1]}
    }

    return nil
}
