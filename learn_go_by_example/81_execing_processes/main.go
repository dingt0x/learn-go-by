package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
    "time"
)

func main() {

    go func() {
        binary, lookErr := exec.LookPath("ls")
        fmt.Println(binary, lookErr)

        args := []string{"ls", "-l", "-a", "-h"}

        env := os.Environ()
        execErr := syscall.Exec(binary, args, env)
        fmt.Println("hi")

        fmt.Println("exec err:", execErr)

    }()

    time.Sleep(time.Second * 8)

    fmt.Println("hi")

}
