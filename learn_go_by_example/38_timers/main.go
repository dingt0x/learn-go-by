package main

import (
    "fmt"
    "time"
)

func do2() {
    t := time.NewTimer(time.Second)

    go func() {
        time.Sleep(time.Millisecond * 500)
        t.Stop()
    }()

    fmt.Println("hi 1")
    fmt.Println("hi 2")
    <-t.C
    fmt.Println("hi 3")
    fmt.Println("hi 4")

}

func do() {
    t1 := time.NewTimer(time.Second * 2)
    <-t1.C
    fmt.Println("Timer 1 fired")

    t2 := time.NewTimer(time.Second)

    fmt.Println("go", time.Now())
    go func() {
        defer func() {
            fmt.Println("hi")
        }()
        <-t2.C
        fmt.Println("Timer 2 fired", time.Now())
    }()
    // 如果将timer.Stop() ，那么 <-timer.C 将会一直阻塞哦~
    if t2.Stop() {
        fmt.Println("t2 stopped")
    }
    time.Sleep(time.Millisecond * 1500)

    // fmt.Println("to reset t2")
    // t2.Reset(time.Millisecond * 200)
    // fmt.Println("ok, let's wait")
    // time.Sleep(time.Millisecond * 1500)

}

func main() {
    do()
    // do2()
}
