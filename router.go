package main

import "github.com/gorilla/mux"

func router() *mux.Router {

	routePaths := []string{
		"/account",
		"/new",
	}
	r := mux.NewRouter()
	r.HandleFunc("/", landingPage)
	for _, route := range routePaths {
		r.HandleFunc(route, landingPage)
	}
	return r
}

func (m *mockServer) controlRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/addRouter/{port}/{handle}", m.addRouter)
	r.HandleFunc("/killRouter/{port}", m.killRouter)
	r.HandleFunc("/startRouter/{port}", m.startRouter)
	return r
}
