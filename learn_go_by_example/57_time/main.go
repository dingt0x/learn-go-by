package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println
    now := time.Now()

    p(now)

    then := time.Date(2009, 11, 17, 20, 34, 58, now.Nanosecond(), time.Local)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Location())

    p(then.Weekday())
    p(then.ISOWeek())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))
    p(then.UTC())
    p(then.Local())

    diff := now.Sub(then)
    p(diff)
    p(diff.Hours())
    p(diff.Seconds())
    p(diff.Nanoseconds())
    p(then.Add(diff))
    p(now.Equal(then.Add(diff)))
}
