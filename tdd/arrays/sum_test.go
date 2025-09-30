package main

import (
    "reflect"
    "testing"
)

func TestSum(t *testing.T) {

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}
        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got '%d' but want '%d' given, %v", got, want, numbers)
        }

    })

}

func TestSumAll(t *testing.T) {

    t.Run("many slices", func(t *testing.T) {
        args := [][]int{{1, 2}, {0, 9}}
        want := []int{3, 9}
        // want := "bo"

        got := SumAll(args...)

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got '%v' but want '%v' given, %v", got, want, args)
        }

    })

}

func TestSumAllTails(t *testing.T) {

    checkSums := func(t testing.TB, got, want []int, args [][]int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got '%v' but want '%v' given, %v", got, want, args)
        }
    }
    t.Run("numbers", func(t *testing.T) {
        args := [][]int{{1, 2}, {0, 9}}
        want := []int{2, 9}
        got := SumAllTails(args...)
        checkSums(t, got, want, args)
    })

    t.Run("empty numbers", func(t *testing.T) {
        args := [][]int{{}, {0, 2, 3}}
        want := []int{0, 5}
        got := SumAllTails(args...)
        checkSums(t, got, want, args)

    })
    t.Run("1 item in numbers", func(t *testing.T) {
        args := [][]int{{1}, {0, 4, 5}}
        want := []int{0, 9}
        got := SumAllTails(args...)
        checkSums(t, got, want, args)

    })
}
