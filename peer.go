package main

import (
	// "fmt"

	"net"
)

type Peer struct {
	conn  net.Conn
	msgCh chan []byte
}

func NewPeer(conn net.Conn, msgCh chan []byte) *Peer {
	return &Peer{conn: conn, msgCh: msgCh}
}

func (p *Peer) readLoop() error {
	defer p.conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			return err
		}
		// fmt.Println(string(buf[:n]))
		msgBuff := make([]byte, n)
		copy(msgBuff, buf[:n])
	}
}
