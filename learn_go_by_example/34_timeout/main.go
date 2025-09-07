package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)

    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "hi c1"
    }()

    select {
    case msg1 := <-c1:
        fmt.Println(msg1)
    case <-time.After(time.Second):
        fmt.Println("Timeout for msg1")
    }
    fmt.Println("let's start c2")

    c2 := make(chan string)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "hi c2"

    }()

    select {
    case msg2 := <-c2:
        fmt.Println(msg2)
    case <-time.After(time.Second * 3):
        fmt.Println("Timeout for msg2")
    }

    fmt.Println("Game Over, Bye~")

}
