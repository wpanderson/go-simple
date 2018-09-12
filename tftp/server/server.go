package server

import (
	"container/list"
	"fmt"
	"net"
	"time"
)

// Helper function for logging packet read/write/error output. Logs are added to "tftp-log.log" in the logs directory.
func logger() {

}

// TODO: implement file storage.
type Server struct {
	listenPort    string                // Port of server.
	serverAddress string                // Address of server.
	files         map[string]*list.List // map of files on the server.
}

// Create and return a server object to maintain port and desired address.
func CreateServer(port, address *string) *Server {
	return &Server{
		listenPort:    *port,
		serverAddress: *address,
	}
}

func Test() {
	fmt.Println("I'm in server. Halp.")
}

// Start tftp server by resolving address and listening for a new connection.
func (server *Server) StartServer(port, address *string) error {

	// Resolve UDP address from local host and port.
	addr, error := net.ResolveUDPAddr("udp", *address+":"+*port)
	if error != nil {
		fmt.Println("Could not resolve network address.", error)
		return error
	}
	fmt.Println(addr)

	// Try to listen on UDP
	conn, error := net.ListenUDP("udp", addr)
	if error != nil {
		fmt.Println("Could not listen on on UDP.", error)
		return error
	}

	// Always remember to close a connection "defer" means code will run after
	// surrounding function closes.
	defer conn.Close()

	// Server loop
	// Use an infinite loop to continually check for a request. When one is found
	// grab it and background it on another port so server can still serve other
	// requests.
	for {
		time.Sleep(5000)
		fmt.Println("Server Going....")
	}
	return error
}
