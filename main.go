package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"mime"
)

// windows sometimes has unexpected mimetypes stored in registry. To avoid, we associate often used extensions with expected mimetypes directly.
// REF: https://stackoverflow.com/questions/54835510/getting-mime-type-text-plain-error-in-golang-while-serving-css, 
func FixMimeTypes() {
    err1 := mime.AddExtensionType(".js", "text/javascript")
    if err1 != nil {
        log.Printf("Error in mime js %s", err1.Error())
    }

    err2 := mime.AddExtensionType(".css", "text/css")
    if err2 != nil {
        log.Printf("Error in mime js %s", err2.Error())
    }
}

func main() {
	var handle_dir = "./"
	var port = "9000"
	if len(os.Args) == 2 {
		handle_dir = os.Args[1]
	}
	
	if len(os.Args) == 3 {
		port = os.Args[2]
	}
	
	port = ":" + port

	fmt.Println("serving " + handle_dir + " " + port)

	FixMimeTypes()
	http.Handle("/", http.FileServer(http.Dir(handle_dir)))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
