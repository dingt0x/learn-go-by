package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {
    p := filepath.Join("dir1", "/dir2", "dir4")
    fmt.Println(p)

    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    fmt.Println(filepath.Dir(p))
    fmt.Println(filepath.Base(p))
    fmt.Println(filepath.Base(p))
    fmt.Println(filepath.IsAbs(p))
    fmt.Println(filepath.IsAbs(filepath.Join("/", p)))

    filename := "config.json"

    ext := filepath.Ext(filename)
    fmt.Println(ext)
    fmt.Println(strings.TrimSuffix(filename, ext))

    rel, err := filepath.Rel("a/b", "a/b/t/tilf")

    fmt.Println(rel, err)
    rel2, err := filepath.Rel("ab/b", "/a/b/t/tilf")
    fmt.Println(rel2, err)

}
