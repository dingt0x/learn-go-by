package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    const s = "สวัสดี"
    fmt.Println("Len s:", len(s))
    r := []rune(s)
    fmt.Println("Len r:", len(r))

    const s2 = "hi你好"
    fmt.Println("Len s2:", len(s2))
    r2 := []rune(s2)
    fmt.Println("Len r2:", len(r2))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()
    for i := 0; i < len(s2); i++ {
        fmt.Printf("%x ", s2[i])
    }
    fmt.Println()

    fmt.Println("Rune count of s:", utf8.RuneCountInString(s))
    fmt.Println("Rune count of s2:", utf8.RuneCountInString(s2))

    for idx, runeVal := range s {
        fmt.Printf("%#U starts at %d\n", runeVal, idx)
    }

    for i, j := 0, 0; i < len(s); i += j {
        val, w := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U x starts at %d\n", val, i)
        j = w
        examineRune(val)
    }

    // fmt.Println("\nUsing DecodeRuneInString")
    // for i, w := 0, 0; i < len(s); i += w {
    //     runeValue, width := utf8.DecodeRuneInString(s[i:])
    //     fmt.Printf("%#U starts at %d\n", runeValue, i)
    //     w = width
    //     examineRune(runeValue)
    // }
}
func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
