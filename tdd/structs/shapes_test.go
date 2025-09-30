package main

import (
    "math"
    "testing"
)

func checkArea(t testing.TB, shape Shape, want float64) {

    t.Helper()

    got := shape.Area()
    if got != want {
        t.Errorf("got %g want %g", got, want)
    }

}
func TestPerimeter(t *testing.T) {

    t.Run("rectangle", func(t *testing.T) {
        rectangle := Rectangle{10, 10}
        got := Perimeter(rectangle)
        want := 40.0
        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    })
}

func TestArea(t *testing.T) {

    areaTests := []struct {
        name    string
        shape   Shape
        hasArea float64
    }{
        {
            name:    "rectangle",
            shape:   Rectangle{10.0, 10.0},
            hasArea: 100.0,
        },
        {
            name:    "circle",
            shape:   Circle{10},
            hasArea: math.Pi * 10 * 10,
        },
        {
            name:    "triangle",
            shape:   Triangle{12, 6},
            hasArea: 36.0,
        },
    }

    for _, tt := range areaTests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
            }
        })
    }

    // t.Run("rectangle", func(t *testing.T) {
    //     rectangle := Rectangle{10, 10}
    //     want := 100.0
    //     checkArea(t, rectangle, want)
    //
    // })
    //
    // t.Run("circles", func(t *testing.T) {
    //     circle := Circle{10}
    //     want := 314.1592653589793
    //     checkArea(t, circle, want)
    // })

}
