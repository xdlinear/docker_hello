package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

func main() {
	http.HandleFunc("/", say_hello)
	log.Println("http server listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func say_hello(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    fmt.Fprintln(w, "hello world\n")
    env := os.Environ()
    for _, v := range env {
        fmt.Fprintf(w, "%s\n", v)
    }
    log.Println(r.RemoteAddr, r.Header)
}
