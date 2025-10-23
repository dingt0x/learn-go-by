package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

const (
    finalWord      = "Go!"
    countDownStart = 3
)

type Sleeper interface {
    Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countDownStart; i > 0; i-- {
        fmt.Fprintln(out, i)
        sleeper.Sleep()
    }
    fmt.Fprint(out, finalWord)

}

type ConfigurableSleeper struct {
    duration time.Duration
    sleep    func(duration time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
    s.sleep(s.duration)
}

func main() {
    sleeper := &ConfigurableSleeper{time.Second * 1, time.Sleep}

    Countdown(os.Stdout, sleeper)
}
