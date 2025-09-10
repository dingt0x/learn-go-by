package main

import (
    "fmt"
    "math/rand/v2"
)

func main() {
    fmt.Printf("%d, ", rand.IntN(100))
    fmt.Println(rand.IntN(100))
    fmt.Println(rand.Int())

    fmt.Println(rand.Float64())
    fmt.Println((rand.Float64() * 5) + 5)
    fmt.Println((rand.Float64() * 5) + 5)

    s2 := rand.NewPCG(42, 1024)
    r2 := rand.New(s2)

    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))
    fmt.Println()

    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)

    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()

}
