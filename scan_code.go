package wsii

import "strings"

func ScanLine(line string) string {
    line = strings.TrimSpace(line)

    if !strings.HasPrefix(line, "#include") {
        return ""
    }

    line = line[len("#include"):]

    quoted := substrEnclosedBy(line, '"', '"')

    if quoted != line {
        return quoted
    }

    angled := substrEnclosedBy(line, '<', '>')

    if angled != line {
        return angled
    }

    return ""
}
