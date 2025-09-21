package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, err := os.ReadFile("/tmp/a.txt")
    check(err)
    fmt.Println(string(data))

    // data2, err := os.ReadFile("/tmp/file_not_exists")
    // check(err)
    // fmt.Println(string(data2))

    f, err := os.Open("/tmp/a")
    defer f.Close()
    check(err)

    b1 := make([]byte, 3)
    n1, err := f.Read(b1)
    check(err)

    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
    fmt.Println(string(b1))

    n12, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n12, string(b1[:n12]))
    // n3, err := f.Write([]byte("hi,hello"))
    // fmt.Println(n3, err)
    fmt.Println("--")

    o2, err := f.Seek(1, io.SeekStart)
    check(err)
    fmt.Println(o2)

    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)

    fmt.Printf("%d bytes @ %d\n", n2, o2)
    fmt.Printf("%d bytes: %s\n", n2, string(b2[:n2]))

    b22, err := io.ReadAll(f)
    check(err)
    fmt.Println(string(b22))

    _, err = f.Seek(0, io.SeekStart)
    check(err)

    fmt.Println("---")
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Println(n3)
    fmt.Println(string(b3))
    // b31 := make([]byte, 4)
    // n31, err := io.ReadAtLeast(f, b31, 4)
    // check(err)
    // fmt.Println(n31)

    _, err = f.Seek(0, io.SeekStart)
    check(err)

    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(3)
    check(err)
    fmt.Printf("3 bytes : %s\n", b4)

    _, err = f.Seek(0, io.SeekStart)
    check(err)

    r5 := bufio.NewReader(f)

    r6 := bufio.NewReader(r5)

    b6, err := r6.Peek(2)
    check(err)
    fmt.Println(b6)

    _, err = f.Seek(0, io.SeekStart)
    check(err)

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

}
