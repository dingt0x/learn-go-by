package main

import (
    "os"
    "text/template"
)

func main() {
    t1 := template.New("t1")

    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }

    t1 = template.Must(t1.Parse("Value is {{.}}\n"))

    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 2)
    t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    t2 := Create("t2", "Name: {{.Name}}\n")

    t2.Execute(os.Stdout, struct {
        Name string
    }{Name: "ser"})
    t2.Execute(os.Stdout, map[string]string{"Name": "ad"})

    t3 := Create("t3", "{{if . -}} yes {{ else -}} no {{ end}} \n")

    t3.Execute(os.Stdout, true)
    t3.Execute(os.Stdout, false)

    t4 := Create("t4", "Range: |{{ range . }} {{ . }} | {{ end }}\n")

    t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
    // go func() {
    //     fmt.Println("Panic goroutine started")
    //     time.Sleep(100 * time.Millisecond)
    //     panic("this goroutine will die alone!") // 没有 recover
    // }()
    // for i := 0; i < 5; i++ {
    //     fmt.Printf("Main goroutine: %d\n", i)
    //     time.Sleep(1 * time.Second)
    // }
    //
    // fmt.Println("Main goroutine completed successfully")
    // time.Sleep(5 * time.Second)
    // fmt.Println("Process is still running!")

}
