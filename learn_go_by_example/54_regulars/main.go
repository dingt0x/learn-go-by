package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")
    fmt.Println(r.MatchString("peach"))

    resFind1 := r.FindString("x peach punch")
    resFind2 := r.FindString("x pe f")

    fmt.Println(resFind1)
    fmt.Println(resFind2 == "")

    resFindStringIndex1 := r.FindStringIndex("x peach punch")
    resFindStringIndex2 := r.FindStringIndex("x peac punc")
    fmt.Println(resFindStringIndex1)
    fmt.Println(resFindStringIndex2)

    fmt.Println(r.FindStringSubmatch("peach punch"))
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    fmt.Println(r.FindAllString("peach punch pinch", -1))
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.FindAllStringSubmatch("peach punch", -1))
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch", -1))

    fmt.Println(r.Match([]byte("peach")))

    fmt.Println("regexp", r)
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))

}
