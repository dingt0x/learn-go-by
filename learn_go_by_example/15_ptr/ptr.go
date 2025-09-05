package main

import "fmt"

func zeroVal(iVal int) {
    iVal = 0
}

func zeroPtr(iPtr *int) {
    *iPtr = 0
}

func main() {

    i := 1
    j := 1
    fmt.Printf("Initial, i: %d, j: %d \n", i, j)

    zeroVal(i)
    zeroPtr(&j)

    fmt.Printf("zeroVal to i: %d, zeroPtr to j: %d \n", i, j)

}
