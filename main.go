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

	fmt.Println("serving " + handle_dir)

	http.Handle("/", http.FileServer(http.Dir(handle_dir)))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
