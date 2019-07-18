package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (m *mockServer) mockHandler(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, string("Easy JSON Server Running :"+r.Host+r.URL.Path))

	portPos := strings.Index(r.Host, ":")
	port := r.Host[portPos+1:]
	handle := r.URL.Path

	response := m.handleData[port+":"+handle]
	fmt.Fprintf(w, string(response))

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

func (m *mockServer) postJSON(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	port := params["port"]
	handle := "/" + params["handle"]

	// var inputs map[string]interface{}

	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&inputs); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer r.Body.Close()

	// res, err := json.Marshal(input)

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()
	m.handleData[port+":"+handle] = s

	fmt.Fprintf(w, string("Post data to port handler "+port+handle))
}
