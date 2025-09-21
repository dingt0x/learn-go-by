package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    d1 := []byte("Hello\ngo\n")

    err := os.WriteFile("/tmp/dat1", d1, 0644)
    check(err)

    s, err := os.ReadFile("/tmp/dat1")
    check(err)
    fmt.Println(string(s))
    fmt.Println(strings.Repeat("-", 8))

    f, err := os.Create("/tmp/dat2")
    defer f.Close()
    n2, err := f.Write([]byte{115, 111, 109, 101, 10})

    check(err)
    fmt.Println(n2)

    s2, err := os.ReadFile("/tmp/dat2")
    check(err)
    fmt.Println(string(s2))
    fmt.Println(strings.Repeat("-", 8))

    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Println(n3)

    s3, err := os.ReadFile("/tmp/dat2")
    check(err)
    fmt.Println(string(s3))
    fmt.Println(strings.Repeat("-", 8))

    w := bufio.NewWriter(f)

    n4, err := w.WriteString("buffered\n")

    check(err)

    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()

}
