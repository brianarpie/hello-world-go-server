package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

type helloWorldResponse struct {
  Message string `json:"message"`
}

type helloWorldRequest struct {
  Name string `json:"name"` // `` struct field tags
}

func main() {
  port := 8080

  imagesHandler := http.FileServer(http.Dir("./images"))

  http.HandleFunc("/helloworld", helloWorldHandler)
  http.Handle("/images/", http.StripPrefix("/images/", imagesHandler))

  log.Printf("Server starting on port %v\n", 8080)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

  var request helloWorldRequest
  decoder := json.NewDecoder(r.Body)

  err := decoder.Decode(&request)
  if err != nil {
    http.Error(w, "Bad Request", http.StatusBadRequest)
    return
  }

  response := helloWorldResponse{Message: "Hello " + request.Name}

  encoder := json.NewEncoder(w)
  encoder.Encode(&response)
}

