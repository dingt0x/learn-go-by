package main

import (
    "fmt"
    "sync"
)

type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}
func main() {

    var c = Container{
        counters: make(map[string]int),
    }

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
        for range n {
            c.inc(name)
        }
        wg.Done()
    }

    wg.Add(3)
    go doIncrement("a", 1000)
    go doIncrement("a", 1000)
    go doIncrement("b", 3000)
    wg.Wait()

    fmt.Println(c.counters)

}
