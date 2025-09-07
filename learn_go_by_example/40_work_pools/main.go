package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    fmt.Printf("Worker %d is working", id)
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
    fmt.Printf("Worker %d off line", id)
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for i := range 3 {
        i++
        go worker(i, jobs, results)
    }

    go func() {
        for i := range numJobs {
            i++
            jobs <- i
        }
        close(jobs)
    }()

    for range numJobs {
        fmt.Println(<-results)
    }
}
