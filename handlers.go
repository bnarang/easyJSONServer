package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func mockLandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Easy JSON Server Running"))
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Test successful"))
}

func (m *mockServer) addRouter(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	port := params["port"]
	handle := "/" + params["handle"]

	m.handles[port] = append(m.handles[port], handle)
	// m.handles = append(m.handles, "/test1")
	fmt.Fprintf(w, string("Router Added"))

}

func (m *mockServer) killRouter(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	port := params["port"]
	go func() { m.stop <- port }()
	fmt.Fprintf(w, string("Killing router "+port))
}

func (m *mockServer) startRouter(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	port := params["port"]

	go func() { m.start <- port }()
	fmt.Fprintf(w, string("Start Router"+port))
}
