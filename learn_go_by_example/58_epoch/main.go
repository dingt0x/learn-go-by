package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    fmt.Println(now)

    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    fmt.Println(time.Unix(now.Unix(), int64(now.Nanosecond())))
    fmt.Println(time.Unix(0, now.UnixNano()))
}
