package main

import (
    "fmt"
)

// var arrayName [size]dataType

type Ss struct {
    s []int
}

func main() {

    // Declare An Array
    var a [5]int
    fmt.Println("emp:", a)
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    // Declare and initialize an array
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl-b:", b)
    b = [...]int{11, 12, 13, 14, 15}
    fmt.Println("dcl-b:", b)

    c := [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl-c:", c)

    d := [...]int{8: 10}
    fmt.Println("dcl-d", d)

    // Incorrect : type mismatch
    // var x int64 = intNumbers3[3]
    // Bad Practice: Type annotation is not needed
    // var x int = intNumbers3[3]
    // fmt.Println(x)

    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }

    }
    fmt.Println("2d: ", twoD)
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)

}
