package main

import (
    "fmt"
    "time"
)

func t1() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println(j)
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()
    for j := range 3 {
        j += 1
        time.Sleep(time.Millisecond * 400)
        jobs <- j
    }

    close(jobs)
    <-done
    fmt.Println("Bye")
}

func t2() {
    // 一个chan如果被关闭了
    // 1. 不能往里发送数据，会被panic
    // 2. 正在接收阻塞的channel，如果被关闭，也会panic~
    jobs := make(chan int, 1)
    jobs <- 1

    go func() {
        fmt.Println("Sending data")
        jobs <- 2
    }()

    time.Sleep(time.Millisecond * 400)

    close(jobs)
    // 不能往里发送数据
    // jobs <- 2

    for range 3 {
        time.Sleep(time.Millisecond * 400)
        j, ok := <-jobs
        fmt.Println(j, ok)
    }
}

func main() {
    t1()
    // t2()

}
