package wsii

import (
    "testing"
)

func AssertIncludes(t *testing.T, got, want []string) {
    t.Helper()

    union := make(map[string]int)

    for _, inc := range got {
        union[inc]++
    }

    for _, inc := range want {
        union[inc]--
    }

    var didntFind, spurious []string
    for inc, membership := range union {
        if membership == -1 {
            didntFind = append(didntFind, inc)
        }

        if membership == 1 {
            spurious = append(spurious, inc)
        }
    }

    if didntFind != nil {
        for _, inc := range didntFind {
            t.Errorf("didn't find header: %q", inc)
        }
    }

    if spurious != nil {
        for _, inc := range spurious {
            t.Errorf("spuriously found: %q", inc)
        }
    }
}
