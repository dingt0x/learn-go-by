package main

import (
    "fmt"
    "strconv"
)

func main() {
    f, err := strconv.ParseFloat("2.233", 64)
    fmt.Println(f, err)

    i, err := strconv.ParseInt("123", 0, 64)
    fmt.Println(i, err)

    d, err := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d, err)

    u, err := strconv.ParseUint("823", 0, 64)
    fmt.Println(u, err)

    k, err := strconv.Atoi("233")
    fmt.Println(k, err)
    xi, err := strconv.ParseInt("111", 16, 0)
    fmt.Println(xi, err)

}
