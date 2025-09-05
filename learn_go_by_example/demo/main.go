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
    var intNumbers [8]int
    // Declare and initialize an array
    var intNumbers2 [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
    intNumbers3 := [3]int{}
    fmt.Println(intNumbers)
    fmt.Println(intNumbers2)
    fmt.Println(intNumbers3)

    // Incorrect : type mismatch
    // var x int64 = intNumbers3[3]
    // Bad Practice: Type annotation is not needed
    // var x int = intNumbers3[3]
    // fmt.Println(x)

    var s Ss
    var s2 = new(Ss)

    fmt.Println(nil == s.s)
    fmt.Println(nil == s2)

    var s3 = new([]int)
    fmt.Println(s3 == nil)
    fmt.Println(nil == *s3)

}
