package main

import (
	"log"
)

var R *Relay

func main() {
	log.Println("relay")
	R = NewRelay()
	R.Run()
}
