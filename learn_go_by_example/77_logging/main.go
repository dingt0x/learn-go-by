package main

import (
    "bytes"
    "fmt"
    "log"
    "log/slog"
    "os"
)

func main() {
    log.Println("standard logger")

    log.SetFlags(log.LstdFlags | log.Lmicroseconds)

    log.Println("with Micro")

    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")

    log.SetPrefix("System: ")
    log.Println("set prefix")

    mylog := log.New(os.Stdout, "my:", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

    mylog.SetOutput(os.Stderr)
    mylog.Println("from mylog")
    mylog.SetPrefix("ohmy:")
    mylog.Println("from mylog")

    var buf bytes.Buffer

    bufLog := log.New(&buf, "buf:", log.LstdFlags)

    bufLog.Println("hi")

    fmt.Println("from buglog:", buf.String())
    bufLog.Println("hello")

    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)

    myslog.Info("hi here")
    myslog.Info("hello again", "key", "val", "age", 25)
}
