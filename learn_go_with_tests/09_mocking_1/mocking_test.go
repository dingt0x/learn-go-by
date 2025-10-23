package main

import (
    "reflect"
    "testing"
)

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
    s.Calls = append(s.Calls, string(p))
    return
}
func TestCountdown(t *testing.T) {

    // buffer := &bytes.Buffer{}
    // spySleeper := &SpySleeper{}
    spySleeperPrinter := &SpyCountdownOperations{}

    countdown(spySleeperPrinter, spySleeperPrinter)

    got := spySleeperPrinter.Calls
    want := []string{
        write,
        "3\n",
        sleep,
        write,
        "2\n",
        sleep,
        write,
        "1\n",
        sleep,
        write,
        "Go!",
    }

    if !reflect.DeepEqual(want, got) {
        t.Errorf("wanted calls %v, got %v", want, got)
    }

}
