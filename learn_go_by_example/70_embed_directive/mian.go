package main

import (
    "embed"
    "fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
    fmt.Println(fileString)
    fmt.Println(fileByte)

    c, err := folder.ReadDir("folder")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(c)
    }

}
