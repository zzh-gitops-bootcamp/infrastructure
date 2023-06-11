package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, World!")
    })
    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, User!")
    })
    http.HandleFunc("/friend", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, Friend!")
    })
    http.ListenAndServe(":8080", nil)
}