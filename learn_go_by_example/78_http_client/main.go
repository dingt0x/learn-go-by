package main

import (
    "bufio"
    "context"
    "fmt"
    "net/http"
    "strings"
    "time"
)

func Get(url string) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://gobyexample.com", nil)
    if err != nil {
        panic(err)
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil || resp == nil {
        panic(err)
    }

    defer resp.Body.Close()

    fmt.Println("response 2 status: ", resp.Status, resp.StatusCode)
    scanner := bufio.NewScanner(resp.Body)
    for range 5 {
        if scanner.Scan() {
            fmt.Println(scanner.Text())
        } else {
            break
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

}

func main() {
    url := "https://gobyexample.com"
    resp, err := http.Get(url)
    if err != nil {
        panic(err)

    }

    defer resp.Body.Close()

    fmt.Println("response status: ", resp.Status, resp.StatusCode)

    scanner := bufio.NewScanner(resp.Body)
    for range 5 {
        if scanner.Scan() {
            fmt.Println(scanner.Text())
        } else {
            break
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println(strings.Repeat("------", 5))

    Get(url)

}
