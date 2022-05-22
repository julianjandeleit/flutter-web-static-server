package main

import (
"net/http"
"log"
)

func main() {
http.Handle("/", http.FileServer(http.Dir("./")))
        err := http.ListenAndServe(":9000", nil)
        if err != nil {
                log.Fatal(err)
        }
}
