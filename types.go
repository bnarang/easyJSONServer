package main

import "net/http"

type mockServer struct {
	start     chan bool
	stop      chan bool
	serveInfo *http.Server
	handles   []string
}

type handler struct {
	http.HandlerFunc
}

type handlers map[string]*handler
