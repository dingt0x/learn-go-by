package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Millisecond * 250)
        c1 <- "hi c1"
    }()
    go func() {
        time.Sleep(time.Millisecond * 250)
        c2 <- "hi c2"
    }()

    for range 2 {
        select {
        case msg1 := <-c1:
            fmt.Println(msg1)
        case msg2 := <-c2:
            fmt.Println(msg2)
        case t := <-time.After(time.Second):
            fmt.Println("time out:", t)
        }
    }
}
