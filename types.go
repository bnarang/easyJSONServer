package main

import "net/http"

type mockServer struct {
	start     chan string
	stop      chan string
	serveInfo map[string]*http.Server
	handles   map[string][]string
}

type handler struct {
	http.HandlerFunc
}

type handlers map[string]*handler
