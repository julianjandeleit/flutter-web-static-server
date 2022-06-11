package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var handle_dir = "./"
	if len(os.Args) == 2 {
		handle_dir = os.Args[1]
	}

	port := ":9000"
	fmt.Println("serving " + handle_dir + " " + port)

	http.Handle("/", http.FileServer(http.Dir(handle_dir)))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
