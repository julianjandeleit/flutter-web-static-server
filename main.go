package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
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
		log.Printf("Error in mime css %s", err2.Error())
	}
}

// from https://stackoverflow.com/questions/39507065/enable-cors-in-golang
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // change this later
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "credentialless")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}

		next.ServeHTTP(w, r)
	})
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
	handle := http.FileServer(http.Dir(handle_dir))
	handle = CORSMiddleware(handle)
	http.Handle("/", handle)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
