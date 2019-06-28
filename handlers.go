package main

import (
	"fmt"
	"net/http"
)

func mockLandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Easy JSON Server Running"))
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Test successful"))
}

func addRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Add a router"))
}

func (m *mockServer) killRouter(w http.ResponseWriter, r *http.Request) {

	go func() { m.stop <- true }()
	fmt.Fprintf(w, string("Killing router"))
}

func (m *mockServer) startRouter(w http.ResponseWriter, r *http.Request) {

	go func() { m.start <- true }()
	fmt.Fprintf(w, string("Start Router"))
}
