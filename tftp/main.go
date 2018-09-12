package main

// Author: Wes Anderson
// Desciprion:
// Start of go-simple a tftp server written in golang. Application starts a simple tftp server which offers in memory tftp servicies.

import (
	"flag"
	"fmt"

	// Server contains logic for starting and maintaining the tftp server.
	"go-simple/tftp/server"
)

// Start in-memory tftp server. Check whether user has supplied port/server flags address then start the server.
func main() {
	fmt.Println("Starting go-simple...")

	// Check flag arguments. use -h to see flags prior to execution.
	port := flag.String("port", "69", "Select a port for go-simple to listen on. (Default: 69)")
	serverAddress := flag.String("address", "localhost", "Select an address to listen on. (Default: localhost)")
	verbose := flag.String("v", "false", "Print verbose output from server operations. (Default: false)")
	flag.Parse()

	srv := server.CreateServer(port, serverAddress)
	fmt.Println(srv)
	fmt.Println("Server port: " + *port)
	fmt.Println("Server Address: " + *serverAddress)
	fmt.Println("Verbose Mode: " + *verbose)

	// Start tftp server. If an error occurs indicate this to the user.
	panic(srv.StartServer(port, serverAddress))
}
