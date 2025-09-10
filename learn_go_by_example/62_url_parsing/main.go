package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    username := "user:name"
    password := "pass@word"

    encodedUser := url.QueryEscape(username) // user%3Aname
    encodedPass := url.QueryEscape(password) // pass%40word
    fmt.Println(encodedUser, encodedPass)

    s := fmt.Sprintf("postgres://%s:%s@host.com:5432/path?k=v&k=2#f", encodedUser, encodedPass)

    u, err := url.Parse(s)

    if err != nil {
        panic(err)
    }
    fmt.Println(u)
    fmt.Println(u.User.Username())
    fmt.Println(u.User.Password())
    fmt.Println(u.User.String())

    fmt.Println(u.Host)

    host, port, err := net.SplitHostPort(u.Host)
    fmt.Println(host, port, err)

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)

    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m["k"])
    fmt.Println(m["k"][0])

}
