package integers

import (
    "fmt"
    "testing"
)

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Printf("%d", sum)
    // Output: 6
}

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    excepted := 4

    if sum != excepted {
        t.Errorf("excepted '%d' but got '%d'", excepted, sum)
    }
}
