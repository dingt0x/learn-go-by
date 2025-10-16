package main

import (
    "log"
    "net/http"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
    Greet(w, "world")
}
func main() {
    log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
