package main

import (
    "fmt"
)

type Rect struct {
    width, height int
}

type X interface {
    Area() int
    Perim() int
}

func (r *Rect) Area() int {
    return r.width * r.height
}

func (r Rect) Perim() int {
    return 2*r.height + 2*r.width

}

func main() {
    var rect Rect

    fmt.Println(rect.Area())
    fmt.Println(rect.Perim())

    fmt.Println((&rect).Area())
    fmt.Println((&rect).Perim())

}
