package p2p

// Peer is an interface that represents the remote node
type Peer interface {
}

// transports is anuthing that handles the communication
// between the nodes in the network, this can be of the
// form of TCP, UDP, etc.
type Transport interface {
	ListenAndAccept() error
}
