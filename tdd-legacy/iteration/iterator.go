package iteration

import (
    "fmt"
    "strings"
)

const repeatCount = 5

func Repeat(character string) string {
    var repeated strings.Builder
    for i := 0; i < repeatCount; i++ {
        repeated.WriteString(character)
    }
    return repeated.String()
}

func RepeatV1(character string) string {
    var repeated string
    for range repeatCount {
        repeated += character
    }
    return repeated

}
func RepeatV2(character string) string {
    var repeated string
    for range repeatCount {
        repeated = fmt.Sprintf("%s%s", repeated, character)
    }
    return repeated
}
