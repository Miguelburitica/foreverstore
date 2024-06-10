package p2p

// handshakeFunc is a function that
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
