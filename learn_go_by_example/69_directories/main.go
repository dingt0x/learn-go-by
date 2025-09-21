package main

import (
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
)

func main() {
    d1 := "/tmp/dir1"

    _, err := os.Stat(d1)

    if err != nil {
        err := os.Mkdir("/tmp/dir1", 0755)
        fmt.Println(err)
    } else {
        fmt.Println("File exist")
    }

    // info, err := os.Stat(d1)
    // fmt.Println(info, err)
    //
    // infoNotExist, err := os.Stat("/tmp/notExist")
    // fmt.Println(infoNotExist, err)
    // if err != nil {
    //     if os.IsNotExist(err) {
    //         fmt.Println("Not Exist")
    //     }
    // }

    createEmptyFile := func(name string) {
        d := []byte("")
        err := os.WriteFile(name, d, 0644)
        if err != nil {
            fmt.Println(err)
        }
    }

    createEmptyFile(filepath.Join("/tmp", "a"))

    err = os.MkdirAll(d1, 0755)
    err = os.MkdirAll(filepath.Join("/tmp", "x", "y", "z"), 0755)
    fmt.Println(err)

    c, err := os.ReadDir("/tmp/x")
    fmt.Println(c, err)

    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir(), entry.Type())
    }

    err = os.Chdir("/tmp")
    fmt.Println(err)
    c1, err := os.ReadDir(".")

    for _, entry := range c1 {
        fmt.Println(" ", entry.Name(), entry.IsDir(), entry.Type())
    }

    err = filepath.WalkDir("/tmp/", visit)

}

func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}
