package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func newServer() *mockServer {
	c := &mockServer{}
	c.start = make(chan bool)
	c.stop = make(chan bool)

	return c
}

func (m *mockServer) setupServer() {

	for {
		select {
		case <-m.start:
			fmt.Println("Start the server now")
			var newS = http.Server{
				Addr:         ":4447",
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			}
			router := mux.NewRouter()

			m.handles = append(m.handles, "/")
			for _, h := range m.handles {
				router.HandleFunc(h, mockLandingPage)
			}
			newS.Handler = router
			m.serveInfo = &newS
			go func() {
				err := newS.ListenAndServe()
				if err != nil {
					fmt.Println(err)
				}
			}()
		case <-m.stop:
			fmt.Println("Stop the server now")
			err := m.serveInfo.Shutdown(context.Background())
			fmt.Println("Error is :", err)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
