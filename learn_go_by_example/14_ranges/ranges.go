package main

import "fmt"

func main() {
    m1 := map[string]int{"x": 2, "a": 1, "b": 2, "B": 22, "1a": 3, "A": 4}

    for k := range m1 {
        fmt.Println(k)
    }

    for k, v := range m1 {
        fmt.Println(k, v)
    }

}
