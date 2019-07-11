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

func (m *mockServer) addRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Add a router"))
	m.handles = append(m.handles, "/test1")
	fmt.Println("M Handles are :", m.handles)
	fmt.Fprintf(w, string("Router Added"))

}

func (m *mockServer) killRouter(w http.ResponseWriter, r *http.Request) {

	go func() { m.stop <- true }()
	fmt.Fprintf(w, string("Killing router"))
}

func (m *mockServer) startRouter(w http.ResponseWriter, r *http.Request) {

	go func() { m.start <- true }()
	fmt.Fprintf(w, string("Start Router"))
}
