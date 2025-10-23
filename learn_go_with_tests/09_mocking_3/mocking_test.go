package main

import (
    "bytes"
    "reflect"
    "testing"
    "time"
)

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

type SpyCountdownOperations struct {
    Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *SpyCountdownOperations) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

func TestCountDown(t *testing.T) {
    t.Run("prints 3 to Go!", func(t *testing.T) {
        buffer := &bytes.Buffer{}
        spySleeper := &SpySleeper{}
        Countdown(buffer, spySleeper)
        got := buffer.String()
        want := "3\n2\n1\nGo!"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }

        if spySleeper.Calls != 3 {
            t.Errorf("not engouth call to sleeper, want 3 got %d", spySleeper.Calls)
        }
    })
    t.Run("sleep before every print", func(t *testing.T) {
        spySleeperPrinter := &SpyCountdownOperations{}
        Countdown(spySleeperPrinter, spySleeperPrinter)

        want := []string{
            write, sleep, write, sleep, write, sleep, write,
        }
        got := spySleeperPrinter.Calls

        if !reflect.DeepEqual(want, got) {
            t.Errorf("want calls %v, got %v", want, got)
        }
    })

}

type SpyTime struct {
    durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
    s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second
    spyTime := &SpyTime{}

    sleeper := &ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
    sleeper.Sleep()

    got := spyTime.durationSlept

    if sleepTime != got {
        t.Errorf("should have splet for %v, but slept for %v", sleepTime, got)
    }

}
