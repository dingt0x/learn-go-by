package main

import (
    "flag"
    "fmt"
)

func main() {
    wortPtr := flag.String("word", "foo", "a string")
    numPtr := flag.Int("numb", 42, "an int")
    forkPtf := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()
    fmt.Println("word", *wortPtr)
    fmt.Println("num", *numPtr)
    fmt.Println("fork", *forkPtf)
    fmt.Println("svar", svar)

}
