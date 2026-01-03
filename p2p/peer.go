package p2p

import "net"

// Peer struct
type Peer struct {
	conn net.Conn // peer connection
}

func (p *Peer) Send(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}
