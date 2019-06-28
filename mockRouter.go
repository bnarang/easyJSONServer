package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func newServer() *mockServer {
	c := &mockServer{
		Server: http.Server{
			Addr:         ":4447",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		stop:  make(chan bool),
		start: make(chan bool),
	}

	router := mux.NewRouter()
	router.HandleFunc("/", mockLandingPage)
	c.Handler = router
	return c
}

func (m *mockServer) setupServer() {

	for {
		select {
		case <-m.start:
			fmt.Println("Start the server now")

			go func() {
				err := m.ListenAndServe()
				if err != nil {
					fmt.Println(err)
				}
			}()
		case <-m.stop:
			fmt.Println("Stop the server now")
			err := m.Shutdown(context.Background())
			fmt.Println("Error is :", err)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
