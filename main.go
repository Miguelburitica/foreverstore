package main

import (
	"log"

	"github.com/Miguelburitica/foreverstore/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error listening and accepting: %v", err)
	}
	select {}
}
