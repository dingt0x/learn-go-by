package main

import "fmt"

func main() {

    for i := range 3 {
        fmt.Println(i)
    }
    for i := 8; i <= 10; i++ {
        fmt.Println(i)
    }

    fmt.Println("---")
    for i := 1; i < 10 && i%2 == 0; i++ {
        fmt.Printf("%d%%2=%d \n", i, i%2)
    }
    fmt.Println("---")
}
