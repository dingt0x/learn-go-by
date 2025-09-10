package main

import (
    "fmt"
    "os"
)

type pointer struct {
    x, y int
}

var f = fmt.Printf

func main() {
    p := pointer{1, 2}
    f("struct1:   %v\n", p)
    f("struct2:   %+v\n", p)
    f("struct3:   %#v\n", p)

    f("type:      %T\n", p)
    f("bool:      %t\n", true)
    f("int:       %d\n", 23)
    f("int2:      %03d\n", 23)
    f("bin:       %b\n", 23)
    f("char:      %c\n", 97)
    f("hex:       %x\n", 456)

    f("float1:    %f\n", 23.823)
    f("float1.2:  %8.2f\n", 23.823)
    f("float1.3:  %08.2f\n", 23.823)

    f("float2:    %e\n", 23.823)
    f("float3:    %E\n", 23.823)

    f("str1:      %s\n", "\"string\"")
    f("str2:      %q\n", "\"string\"")
    f("str2.2:    %q\n", "string")
    f("str3:      %x\n", "hex this")

    f("pointer:   %p\n", &p)
    f("width1:    |%6d|%6d|\n", 23, 233)

    f("width2:    |%6.2f|%6.2f|\n", 1.2, 3.23)

    f("width3:    |%-6.2f|%-6.2f|\n", 1.2, 3.23)

    f("width4:    |%6s|%6s|\n", "foo", "b")
    f("width5:    |%-6s|%-6s|\n", "foo", "b")

    fmt.Fprintf(os.Stderr, "ab %s", "sdf")

}
