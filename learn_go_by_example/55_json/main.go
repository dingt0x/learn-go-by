package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    boolB, err := json.Marshal(true)
    fmt.Println(string(boolB), err)

    intB, err := json.Marshal(1)
    fmt.Println(string(intB), err)

    fltB, err := json.Marshal(2.34)
    fmt.Println(string(fltB), err)

    x := "sd"
    strB, err := json.Marshal(x)
    fmt.Println(string(strB), err)

    slcD := []string{"apple", "peach", "pear"}
    slcB, err := json.Marshal(slcD)
    fmt.Println(string(slcB), err)

    mapD := map[string]int{"apple": 4, "lettuce": 3}
    mapB, err := json.Marshal(mapD)
    fmt.Println(string(mapB), err)

    res1D := &response1{1, []string{"apple", "peach", "pear"}}
    res1B, err := json.Marshal(res1D)
    fmt.Println(string(res1B), err)

    res2D := &response2{1, []string{"apple", "peach", "pear"}}
    res2B, err := json.Marshal(res2D)
    fmt.Println(string(res2B), err)

    byt := []byte(`{"num":6.12,"strs":["a","b"]}`)

    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    strs, ok1 := dat["strs"].([]interface{})
    str1, ok := strs[0].(string)
    fmt.Println(ok1)
    fmt.Println(str1, ok)

    str := `{"page":1,"fruits":["apple","peach","pear"]}`
    res := response2{}

    if err := json.Unmarshal([]byte(str), &res); err != nil {
        panic(err)
    }

    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    f, _ := os.Create("/tmp/a.txt")
    enc := json.NewEncoder(f)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)

    dec := json.NewDecoder(strings.NewReader(str))
    res1 := response2{}
    dec.Decode(&res1)
    fmt.Println(res1)

}
