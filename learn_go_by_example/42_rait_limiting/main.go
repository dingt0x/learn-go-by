package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)

    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(1000 * time.Millisecond)
    fmt.Println(time.Now())

    for req := range requests {
        time.Sleep(1500 * time.Millisecond)
        t := <-limiter
        fmt.Println("request:", req, t, time.Now())
    }

    burstLimiter := make(chan time.Time, 3)
    for range 3 {
        burstLimiter <- time.Now()
    }

    go func() {
        ticker := time.NewTicker(time.Millisecond * 500)
        for {
            t := <-ticker.C
            burstLimiter <- t
        }
        // for t := range time.Tick(time.Millisecond * 300) {
        //     burstLimiter <- t
        // }
    }()

    burstRequest := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstRequest <- i
    }

    close(burstRequest)

    for req := range burstRequest {
        <-burstLimiter
        fmt.Println("request", req, time.Now())

    }

}
