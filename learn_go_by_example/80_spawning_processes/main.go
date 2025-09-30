package main

import (
    "fmt"
    "io"
    "os"
    "os/exec"
)

func main() {
    fmt.Println("PATH:", os.Getenv("PATH"))
    dateCmd := exec.Command("date")

    dateOut, err := dateCmd.Output()

    if err != nil {
        panic(err)
    }

    fmt.Println(string(dateOut))

    cmdExec := exec.Command("date", "-sdx")
    cmdOutput, err := cmdExec.Output()

    fmt.Println("cmdOutput: ", cmdOutput)
    if err != nil {

        switch e := err.(type) {
        case *exec.Error:
            fmt.Println("failed executing:", e)
        case *exec.ExitError:
            fmt.Println("*exec.ExitError", e)
            fmt.Println("command exit rc =", e.ExitCode())
        default:
            panic(err)
        }
    }

    grepCmd := exec.Command("grep", "http-world")

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()

    grepCmd.Start()

    grepIn.Write([]byte("http-world grep\ngoodbye grep"))
    grepIn.Close()

    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep http-world")
    fmt.Println(string(grepBytes))

}
