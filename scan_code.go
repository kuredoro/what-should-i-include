package wsii

import (
	"bufio"
	"path/filepath"
	"strings"
)

func ScanLine(line string) string {
    line = strings.TrimSpace(line)

    if !strings.HasPrefix(line, "#include") {
        return ""
    }

    line = line[len("#include"):]

    quoted := substrEnclosedBy(line, '"', '"')

    if quoted != line {
        return filepath.Clean(quoted)
    }

    angled := substrEnclosedBy(line, '<', '>')

    if angled != line {
        return filepath.Clean(angled)
    }

    return ""
}

func ScanCode(code string) []string {
    var headers []string

    buf := strings.NewReader(code)
    s := bufio.NewScanner(buf)

    for s.Scan() {
        header := ScanLine(s.Text())

        if header != "" {
            headers = append(headers, header)
        }
    }

    return headers
}
