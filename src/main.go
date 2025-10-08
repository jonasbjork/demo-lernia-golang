package main

import (
"fmt"
"log"
"net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Lernia!")
  })

  mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "ok")
  })

  log.Println("api: listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", mux))
}

