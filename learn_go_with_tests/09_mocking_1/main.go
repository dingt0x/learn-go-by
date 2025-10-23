package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

const finalWord = "Go!"
const countdownStr = 3

func countdown(w io.Writer, sleeper Sleeper) {
    for i := countdownStr; i > 0; i-- {
        fmt.Fprintln(w, i)
        sleeper.Sleep()

    }
    fmt.Fprintf(w, finalWord)
}

type Sleeper interface {
    Sleep()
}

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

func main() {
    countdown(os.Stdout, &DefaultSleeper{})
}
