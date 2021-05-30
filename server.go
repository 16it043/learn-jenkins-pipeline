package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s %s\n", request.RemoteAddr, request.Method, request.RequestURI)

		response.Header().Set("Content-Type", "text/html")

		fmt.Fprint(response, "<h1>Hello World!</h1>")
	})

	log.Println("Starting the server on 0.0.0.0:80")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
