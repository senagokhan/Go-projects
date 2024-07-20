package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	x1 := &messageHandler{"Bu birinci mesaj!"}
	x2 := &messageHandler{"Bu ikinci mesaj!"}

	mux.Handle("/bir", x1)
	mux.Handle("/iki", x2)

	log.Println("Starting server")
	http.ListenAndServe(":8080", mux)

}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x.message)
}
