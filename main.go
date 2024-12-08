package main

import (
	"fmt"
	"log/slog"
	"net"
)

const defaultListner = ":5050"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	peers     map[*Peer]bool
	ln        net.Listener
	addPeerCh chan *Peer
}

// ** create a new server instance
func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListner
	}
	return &Server{Config: cfg,
		peers:     make(map[*Peer]bool),
		addPeerCh: make(chan *Peer)}
}

// ** start the server bro
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	go s.loop()

	return s.AcceptLoop()
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
		default:
			fmt.Println("foo")
		}
	}
}

func (s *Server) AcceptLoop() error {
	for {

		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "err", err)
			continue
		}

		go s.handleConn(conn)

	}

}

func (s *Server) handleConn(conn net.Conn) {

	peer := NewPeer(conn)
	s.addPeerCh <- peer

}

func main() {

}
