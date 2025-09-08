package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting work\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    fmt.Println("hi")
    var wg sync.WaitGroup
    for i := range 5 {
        wg.Go(func() {
            i++
            worker(i)
        })
    }
    wg.Wait()
}
