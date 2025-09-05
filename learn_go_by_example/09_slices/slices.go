package main

import (
    "fmt"
    "slices"
)

func main() {
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s))

    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len:", len(s))

    s = append(s, "d")
    fmt.Println("apd len:", len(s))
    fmt.Println("apd cap:", cap(s))
    s = append(s, "e", "f")
    fmt.Println("apd-2 len:", len(s))
    fmt.Println("apd-2 cap:", cap(s))

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println(c, s)

    l := s[2:5]
    fmt.Println("sl1", l)
    fmt.Println("sl1 len", len(l))
    fmt.Println("sl1 cap", len(l))

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    fmt.Println(s)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    var s1 []string
    fmt.Println("len of s1", len(s1), "s1 is nil:", s1 == nil)
    var s2 = []string{}
    fmt.Println("len of s2", len(s2), "s2 is nil:", s2 == nil)

    s4 := make([]string, 0, 3)

    fmt.Println("len of s4", len(s4), "s4 is nil:", s4 == nil)

    s9 := make([]int, 0)
    fmt.Println(len(s9))

}
