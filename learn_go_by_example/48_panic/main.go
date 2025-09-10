package main

import (
    "fmt"
    "time"
)

func main() {

    go func() {
        panic("a problem")
    }()

    time.Sleep(time.Second)

    fmt.Println("hi")

}
