package network

import (
	"net"
)

// Serveur gestion connexion
func HandleConnection(conn net.Conn) {
	// Fermeture automatique de la connexion
	defer conn.Close()

	message := "Dock Lift CLI\n"
	// Envoi du message au client
	conn.Write([]byte(message))
}

// Client connexion au serveur
func Connect(address string) (string, error) {
	// Établissement de la connexion TCP
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Lecture de la réponse du serveur
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}
