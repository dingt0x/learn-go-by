package main

import (
    "fmt"
    "time"
)

func main() {
    msgs := make(chan string)

    go func() { msgs <- "hello" }()
    msg := <-msgs
    fmt.Println(msg)

    messages := make(chan string, 2)
    messages <- "buffered"
    messages <- "channel"
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    done := make(chan bool)

    go worker(done)

    <-done
}

func worker(done chan bool) {
    fmt.Println("working...")
    time.Sleep(time.Second * 1)
    fmt.Println("done")

    done <- true
}
