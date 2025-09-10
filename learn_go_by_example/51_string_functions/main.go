package main

import (
    "fmt"
    "strings"
)

var p = fmt.Println

func main() {
    p("Contains:    ", strings.Contains("test", "s"))
    p("Count:       ", strings.Count("test", "t"))
    p("HasPrefix:   ", strings.HasPrefix("test", "te"))
    p("HasSuffix:   ", strings.HasSuffix("test", "st"))
    p("Index:       ", strings.Index("test", "s"))
    p("Join:        ", strings.Join([]string{"a", "b"}, "-"))
    p("Repeat:      ", strings.Repeat("a", 5))
    p("Replace:     ", strings.Replace("test", "t", "T", -1))
    p("Replace:     ", strings.Replace("test", "t", "T", 1))
    p("Split:       ", strings.Split("a-b-c-d-e", "-"))
    p("SplitN:      ", strings.SplitN("a-b-c-d-e", "-", 2))
    p("SplitAfter:  ", strings.SplitAfter("a-b-c", "-"))
    p("SplitAfterN: ", strings.SplitAfterN("a-b-c", "-", 2))
    p("ToLower:     ", strings.ToLower("aBbc"))
    p("ToUpper:     ", strings.ToUpper("test"))
    p("ToTitle:     ", strings.ToTitle("test"))
}
