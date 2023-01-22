package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Привет"))
}

func main() {
    
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
 
    
    log.Println("Запуск веб-сервера на http://localhost:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}