package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
)

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("tpl/index.html")
  if err != nil {
    w.WriteHeader(500)
    return
  }
  err = t.Execute(w, nil)
  if err != nil {
    w.WriteHeader(500)
  }
}

func main() {
  http.HandleFunc("/", index)
  err := http.ListenAndServe("0.0.0.0:" + os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatalf("Error: %s", err)
  }
}
