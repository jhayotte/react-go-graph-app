package main

import (
    "net/http"
    "log"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./frontend/build")))
    
    log.Println("Server starting on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
