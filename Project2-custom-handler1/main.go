package main

import (
	"io"
	"net/http"
)

func main() {
	var i ironman
	var w wolverine

	mux := http.NewServeMux()
	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

type ironman int

func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {

	io.WriteString(res, "Mr. Iron!")
}

type wolverine int

func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Wolverine!")
}
