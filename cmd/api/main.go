package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

var books []Book // объект структуры Book

func main(){
  r := mux.NewRouter()

  setupRoutes(r)
  log.Fatal(http.ListenAndServe(":8080", r))
}
