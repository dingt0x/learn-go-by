package main

import (
    "fmt"
    "os"
)

func main() {
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("Create file: ", p)
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func closeFile(f *os.File) {
    fmt.Println("Closing file:", f.Name())
    err := f.Close()
    if err != nil {
        panic(err)
    }
}

func writeFile(f *os.File) {
    fmt.Println("Writing file:", f.Name())
    fmt.Fprintln(f, "data2")
}
