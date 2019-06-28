package main

import "net/http"

type mockServer struct {
	http.Server
	start chan bool
	stop  chan bool
}
