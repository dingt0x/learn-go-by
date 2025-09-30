package main

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
import (
    "bytes"
    "testing"
)

func TestCountDown(t *testing.T) {
    buffer := &bytes.Buffer{}

    Countdown(buffer)

    got := buffer.String()
    want := "3\n2\n1\nGo!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }

}
