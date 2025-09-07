package main

import (
    "fmt"
)

func main() {
    c1 := make(chan string)

    select {
    case msg1 := <-c1:
        fmt.Println("Message from channel:", msg1)
    default:
        fmt.Println("no message for channel")
    }

    select {
    case c1 <- "message":
        fmt.Println("message send to channel")
    default:
        fmt.Println("No message send to channel")
    }

    signals := make(chan bool)

    select {
    case c1 <- "message":
        fmt.Println("message send to channel")
    case signals <- true:
        fmt.Println("bool send to channel")
    default:
        fmt.Println("no activity")
    }

    // go func() {
    //     for i := range 10 {
    //         time.Sleep(time.Millisecond * 500)
    //         c1 <- fmt.Sprintf("msg-%d", i)
    //     }
    //     close(c1)
    //
    // }()
    //
    // for i := range c1 {
    //     fmt.Println(i)
    //
    // }

}
