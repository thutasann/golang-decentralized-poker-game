package p2p

import (
	"fmt"
	"net"
)

// Server Config struct
type ServerConfig struct {
	ListenAddr string // listen address
}

// Server struct
type Server struct {
	ServerConfig

	listener net.Listener // net listener
	// mu       sync.RWMutex       // mutex
	peers   map[net.Addr]*Peer // peers map
	addPeer chan *Peer         // add peer channel
}

// Initialize a new Server
func NewServer(cfg ServerConfig) *Server {
	return &Server{
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
	}
}

// Start the Server
//
// telnet localhost 3000
func (s *Server) Start() {
	// loop
	go s.loop()

	// listen
	if err := s.listen(); err != nil {
		panic(err)
	}

	fmt.Printf("game server running on port %s\n", s.ListenAddr)

	// accept loop
	s.acceptLoop()
}

// Listen the server
func (s *Server) listen() error {
	listen, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		panic(err)
	}

	s.listener = listen
	return nil
}

// Loop the server
func (s *Server) loop() {
	for peer := range s.addPeer {
		s.peers[peer.conn.RemoteAddr()] = peer
		fmt.Printf("New Player connected: %s\n", peer.conn.RemoteAddr())
	}
}

// Accept the Server Loop
func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}

		peer := &Peer{
			conn: conn,
		}

		s.addPeer <- peer

		peer.Send([]byte("GGPOKER V0.1-beta"))

		go s.handleConn(conn)
	}
}

// Handle the Server Connection
func (s *Server) handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
