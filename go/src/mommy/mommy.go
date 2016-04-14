package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/check/", check)
    http.HandleFunc("/doctor/", call_doctor)
    http.ListenAndServe(":8080", nil)
}

type Message struct {
    Text string
}

func check (w http.ResponseWriter, r *http.Request) {

    m := Message{"Better check on those monkeys"}
    b, err := json.Marshal(m)

    if err != nil {
        panic(err)
    }

     w.Write(b)
}

func call_doctor (w http.ResponseWriter, r *http.Request) {

    m := Message{"Mommy called the doctor"}
    b, err := json.Marshal(m)

    if err != nil {
        panic(err)
    }

     w.Write(b)
}