package main

import (
    "cmp"
    "fmt"
    "slices"
)

func main() {
    strs := []string{"d", "a", "b", "c"}
    slices.Sort(strs)
    fmt.Println("Strings: ", strs)

    ints := []int{2, 8, 3}
    slices.Sort(ints)

    fmt.Println("Ints:    ", ints)

    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)

    fruits := []string{"peach", "banana", "kiwi"}
    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }
    slices.SortFunc(fruits, lenCmp)
    fmt.Println(fruits)

    type Person struct {
        name string
        age  int
    }

    people := []Person{
        {"Jax", 37},
        {"TJ", 24},
        {"Alex", 98},
    }

    slices.SortFunc(people, func(a, b Person) int {
        return cmp.Compare(a.age, b.age)
    })

    fmt.Println(people)

}
