package main

import (
    "fmt"
    "time"
)

func t1() {

    c1 := make(chan int)

    go func() {
        for i := range 10 {
            time.Sleep(time.Millisecond * 500)
            c1 <- i
        }
        close(c1)

    }()

    for e := range c1 {
        fmt.Println("Received data", e)
    }

}

func main() {
    t1()
}
