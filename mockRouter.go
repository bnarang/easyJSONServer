package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func newServer() *mockServer {

	m := &mockServer{}
	m.start = make(chan string)
	m.stop = make(chan string)
	m.handles = make(map[string][]string, 3)
	m.serveInfo = make(map[string]*http.Server, 3)
	m.handleData = make(map[string]string, 3)
	return m
}

func (m *mockServer) setupServer() {

	for {
		select {
		case port := <-m.start:
			fmt.Println("Start the server now")
			var newS = http.Server{
				Addr:         ":" + port,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			}
			router := mux.NewRouter()

			m.handles[port] = append(m.handles[port], "/")
			fmt.Println("Port handlers are :", port, m.handles[port])
			for _, h := range m.handles[port] {
				router.HandleFunc(h, m.mockHandler)
			}
			newS.Handler = router
			m.serveInfo[port] = &newS
			go func() {
				err := newS.ListenAndServe()
				if err != nil {
					fmt.Println(err)
				}
			}()
		case port := <-m.stop:
			fmt.Println("Stop the server now " + port)
			err := m.serveInfo[port].Shutdown(context.Background())
			fmt.Println("Error is :", err)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
