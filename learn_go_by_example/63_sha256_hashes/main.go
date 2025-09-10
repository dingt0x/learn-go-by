package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {

    s := "sha256 this string"
    s2 := "sha256  --- this string"

    s8 := fmt.Sprintf("%s%s", s, s2)
    fmt.Println(s)
    h := sha256.New()
    fmt.Println(h)
    h.Write([]byte(s))
    fmt.Println(h)

    bs1 := h.Sum(nil)
    fmt.Printf("%x\n", bs1)

    h.Write([]byte(s2))
    fmt.Println(h)

    bs := h.Sum(nil)
    fmt.Printf("%x\n", bs)

    h2 := sha256.New()
    h2.Write([]byte(s8))
    bs8 := h.Sum(nil)
    fmt.Printf("%x\n", bs8)

}
