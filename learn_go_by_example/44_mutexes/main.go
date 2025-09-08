package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var mu sync.Mutex

    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        for {
            mu.Lock()
            fmt.Println("HI, 2")
            time.Sleep(time.Millisecond * 210)
            mu.Unlock()
        }
    }()
    go func() {
        for {
            mu.Lock()
            fmt.Println("HI, 1")
            time.Sleep(time.Millisecond * 200)
            mu.Unlock()
        }
    }()

    wg.Wait()

}
