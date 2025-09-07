package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(time.Millisecond * 500)
    done := make(chan struct{})

    go func() {
        for {
            select {
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            case <-done:
                return
            }
        }
    }()
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    done <- struct{}{}

    fmt.Println("Ticker stopped")
}
