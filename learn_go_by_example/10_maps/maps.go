package main

import (
    "fmt"
)

func main() {
    // declare a map
    m1 := make(map[string]int)
    var m2 map[string]int
    fmt.Println(m2 == nil)

    // declare and initialize a map

    m3 := map[string]int{}
    fmt.Println(m3 == nil)

    m1["k1"] = 1
    m1["k2"] = 2

    v1 := m1["k1"]
    fmt.Println(v1)
    _, prs := m1["k2"]
    fmt.Println("prs:", prs)
    _, prsX := m1["kx"]
    fmt.Println("prsX:", prsX)

    // Initialize

    if v2, prs := m1["k2"]; prs {
        fmt.Println("v2:", v2)
    } else {
        fmt.Println("Not found v2")
    }

}
