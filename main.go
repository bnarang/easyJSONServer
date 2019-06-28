package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func landingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Test successful"))
}

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

func addRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Add a router"))
}

func killRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("Killing router"))
}

func controlRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/addRouter", addRouter)
	r.HandleFunc("/killRouter", killRouter)
	return r
}

func (H *hub) routeExecutor() {
	for {
		select {
		case <-H.start:
			log.Fatal(http.ListenAndServe(":4446", router()))
		case <-H.stop:
			return
		}
	}

}
func main() {

	go func() {
		log.Fatal(http.ListenAndServe(":4446", router()))
	}()
	http.ListenAndServe(":4444", controlRouter())
}
