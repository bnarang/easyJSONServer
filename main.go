package main

import (
	"net/http"
)

func main() {

	// go func() {
	// 	log.Fatal(http.ListenAndServe(":4446", router()))
	// }()

	m := newServer()

	go m.setupServer()
	http.ListenAndServe(":4444", m.controlRouter())
}
