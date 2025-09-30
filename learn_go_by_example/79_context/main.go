package main

import (
    "fmt"
    "net/http"
    "time"
)

func hello(w http.ResponseWriter, req *http.Request) {
    ctx := req.Context()

    fmt.Println("Server: http handler started")
    defer fmt.Println("Server: http handler ended")

    select {
    case <-time.After(time.Second * 10):
        fmt.Fprintf(w, "http-world\n")
    case <-ctx.Done():
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}
