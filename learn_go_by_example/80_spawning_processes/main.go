package main

import (
    "fmt"
    "os"
    "os/exec"
)

func returnX() interface{} {
    return "abc"
}
func main() {
    fmt.Println("PATH:", os.Getenv("PATH"))
    dateCmd := exec.Command("date")

    dateOut, err := dateCmd.Output()

    if err != nil {
        panic(err)
    }

    fmt.Println(string(dateOut))

    fmt.Println("---")

    s := returnX()
    switch e := s.(type) {
    case string:
        fmt.Println("Here is a string", e)
    default:
        fmt.Println(e)
    }
    // _, err := exec.Command("date", "-x").Output()
    //
    // if err != nil {
    //     switch e := err.(type) {
    //     case *exec.Error
    //
    //     }
}
