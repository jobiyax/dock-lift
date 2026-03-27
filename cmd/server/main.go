package main

import (
	"fmt"
	"log"
	"net"

	"dock-lift-cli/internal/network"
)

func main() {
	// Écoute sur le port 9090
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	// Fermeture à la fin
	defer listener.Close()

	fmt.Println("🚀 Server running on port 9090...")

	for {
		// Accepter une connexion
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// Gérer la connexion en arrière-plan
		go network.HandleConnection(conn)
	}
}
