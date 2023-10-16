package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello ci-cd!")
  })

  log.Println("Listening on port 6699")
  http.ListenAndServe(":6699", nil)
}
