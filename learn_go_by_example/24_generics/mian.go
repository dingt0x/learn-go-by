package main

import (
    "fmt"
    "iter"
)

type Number interface {
    ~int | ~int64 | ~float64
}
type NNumber interface {
}

func f1(x NNumber) {
    fmt.Println("sd")
}

func SliceIndex[S ~[]E, E comparable](s S, e E) int {
    for k, v := range s {
        if v == e {
            return k
        }
    }
    return -1
}

type List[T any] struct {
    head, tail *element[T]
}
type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {
        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }

}

func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 1, 1
        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

func main() {
    s := []string{"a", "b", "c"}
    fmt.Println(SliceIndex(s, "c"))

    var l List[int]

    l.Push(88)
    l.Push(99)
    fmt.Println(l.AllElements())

    for i := range genFib() {
        fmt.Println(i)
        if i > 10 {
            break
        }
    }

    for v := range l.All() {
        fmt.Println(v)
    }

}
