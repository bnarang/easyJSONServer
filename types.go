package main

type hub struct {
	start chan string
	stop  chan string
}
