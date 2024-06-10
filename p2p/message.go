package p2p

import "net"

// Message holds any artbitrary data that is being sent
// over the each transport between the nodes in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
